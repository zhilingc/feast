import random
import time
from datetime import datetime
from datetime import timedelta
from urllib.parse import urlparse

import numpy as np
import pandas as pd
import pytest
import pytz
from feast.client import Client
from feast.entity import Entity
from feast.feature import Feature
from feast.feature_set import FeatureSet
from feast.type_map import ValueType
from google.cloud import storage
from google.protobuf.duration_pb2 import Duration
from pandavro import to_avro


@pytest.fixture(scope="module")
def core_url(pytestconfig):
    return pytestconfig.getoption("core_url")


@pytest.fixture(scope="module")
def serving_url(pytestconfig):
    return pytestconfig.getoption("serving_url")


@pytest.fixture(scope="module")
def allow_dirty(pytestconfig):
    return True if pytestconfig.getoption("allow_dirty").lower() == "true" else False


@pytest.fixture(scope="module")
def gcs_path(pytestconfig):
    return pytestconfig.getoption("gcs_path")


@pytest.fixture(scope="module")
def client(core_url, serving_url, allow_dirty):
    # Get client for core and serving
    client = Client(core_url=core_url, serving_url=serving_url)
    client.set_project("project1")

    # Ensure Feast core is active, but empty
    if not allow_dirty:
        feature_sets = client.list_feature_sets()
        if len(feature_sets) > 0:
            raise Exception("Feast cannot have existing feature sets registered. Exiting tests.")

    return client


def test_get_batch_features_with_file(client):
    file_fs1 = FeatureSet(
        "file_feature_set",
        features=[Feature("feature_value1", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )

    client.apply(file_fs1)
    file_fs1 = client.get_feature_set(name="file_feature_set", version=1)

    N_ROWS = 10
    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    features_1_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "feature_value1": [f"{i}" for i in range(N_ROWS)],
        }
    )
    client.ingest(file_fs1, features_1_df)

    # Rename column (datetime -> event_timestamp)
    features_1_df = features_1_df.rename(columns={"datetime": "event_timestamp"})

    to_avro(df=features_1_df, file_path_or_buffer="file_feature_set.avro")

    feature_retrieval_job = client.get_batch_features(
        entity_rows="file://file_feature_set.avro", feature_refs=["feature_value1:1"]
    )

    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert output["entity_id"].to_list() == [int(i) for i in output["project1_feature_value1_v1"].to_list()]


def test_get_batch_features_with_gs_path(client, gcs_path):
    gcs_fs1 = FeatureSet(
        "gcs_feature_set",
        features=[Feature("feature_value2", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )

    client.apply(gcs_fs1)
    gcs_fs1 = client.get_feature_set(name="gcs_feature_set", version=1)

    N_ROWS = 10
    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    features_1_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "feature_value2": [f"{i}" for i in range(N_ROWS)],
        }
    )
    client.ingest(gcs_fs1, features_1_df)

    # Rename column (datetime -> event_timestamp)
    features_1_df = features_1_df.rename(columns={"datetime": "event_timestamp"})

    # Output file to local
    file_name = "gcs_feature_set.avro"
    to_avro(df=features_1_df, file_path_or_buffer=file_name)

    uri = urlparse(gcs_path)
    bucket = uri.hostname
    ts = int(time.time())
    remote_path = str(uri.path).strip("/") + f"{ts}/{file_name}"

    # Upload file to gcs
    storage_client = storage.Client(project=None)
    bucket = storage_client.get_bucket(bucket)
    blob = bucket.blob(remote_path)
    blob.upload_from_filename(file_name)

    feature_retrieval_job = client.get_batch_features(
        entity_rows=f"{gcs_path}{ts}/*",
        feature_refs=["feature_value2:1"]
    )

    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert output["entity_id"].to_list() == [int(i) for i in output["project1_feature_value2_v1"].to_list()]


def test_order_by_creation_time(client):
    proc_time_fs = FeatureSet(
        "processing_time",
        features=[Feature("feature_value3", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )
    client.apply(proc_time_fs)
    time.sleep(10)
    proc_time_fs = client.get_feature_set(name="processing_time", version=1)

    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    N_ROWS = 10
    incorrect_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "feature_value3": ["WRONG"] * N_ROWS,
        }
    )
    correct_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "feature_value3": ["CORRECT"] * N_ROWS,
        }
    )
    client.ingest(proc_time_fs, incorrect_df)
    time.sleep(10)
    client.ingest(proc_time_fs, correct_df)
    feature_retrieval_job = client.get_batch_features(
        entity_rows=incorrect_df[["datetime", "entity_id"]], feature_refs=["project1/feature_value3:1"]
    )
    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert output["project1_feature_value3_v1"].to_list() == ["CORRECT"] * N_ROWS


def test_additional_columns_in_entity_table(client):
    add_cols_fs = FeatureSet(
        "additional_columns",
        features=[Feature("feature_value4", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )
    client.apply(add_cols_fs)
    time.sleep(10)
    add_cols_fs = client.get_feature_set(name="additional_columns", version=1)

    N_ROWS = 10
    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    features_df = pd.DataFrame(
        {"datetime": [time_offset] * N_ROWS, "entity_id": [i for i in range(N_ROWS)], "feature_value4": ["abc"] * N_ROWS}
    )
    client.ingest(add_cols_fs, features_df)

    entity_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "additional_string_col": ["hello im extra"] * N_ROWS,
            "additional_float_col": [random.random() for i in range(N_ROWS)],
        }
    )
    feature_retrieval_job = client.get_batch_features(
        entity_rows=entity_df, feature_refs=["project1/feature_value4:1"]
    )
    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert np.allclose(output["additional_float_col"], entity_df["additional_float_col"])
    assert output["additional_string_col"].to_list() == entity_df["additional_string_col"].to_list()
    assert output["project1_feature_value4_v1"].to_list() == features_df["feature_value4"].to_list()


def test_point_in_time_correctness_join(client):
    historical_fs = FeatureSet(
        "historical",
        features=[Feature("feature_value5", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )
    client.apply(historical_fs)
    time.sleep(10)
    historical_fs = client.get_feature_set(name="historical", version=1)

    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    N_EXAMPLES = 10
    historical_df = pd.DataFrame(
        {
            "datetime": [
                time_offset - timedelta(seconds=50),
                time_offset - timedelta(seconds=30),
                time_offset - timedelta(seconds=10),
            ]
            * N_EXAMPLES,
            "entity_id": [i for i in range(N_EXAMPLES) for _ in range(3)],
            "feature_value5": ["WRONG", "WRONG", "CORRECT"] * N_EXAMPLES,
        }
    )
    entity_df = pd.DataFrame(
        {"datetime": [time_offset - timedelta(seconds=10)] * N_EXAMPLES, "entity_id": [i for i in range(N_EXAMPLES)]}
    )

    client.ingest(historical_fs, historical_df)

    feature_retrieval_job = client.get_batch_features(entity_rows=entity_df, feature_refs=["project1/feature_value5:1"])
    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert output["project1_feature_value5"].to_list() == ["CORRECT"] * N_EXAMPLES


def test_multiple_featureset_joins(client):
    fs1 = FeatureSet(
        "feature_set_1",
        features=[Feature("feature_value6", ValueType.STRING)],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )

    fs2 = FeatureSet(
        "feature_set_2",
        features=[Feature("other_feature_value7", ValueType.INT64)],
        entities=[Entity("other_entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )

    client.apply(fs1)
    fs1 = client.get_feature_set(name="feature_set_1", version=1)

    client.apply(fs2)
    fs2 = client.get_feature_set(name="feature_set_2", version=1)

    N_ROWS = 10
    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    features_1_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "feature_value6": [f"{i}" for i in range(N_ROWS)],
        }
    )
    client.ingest(fs1, features_1_df)

    features_2_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "other_entity_id": [i for i in range(N_ROWS)],
            "other_feature_value7": [i for i in range(N_ROWS)],
        }
    )
    client.ingest(fs2, features_2_df)

    entity_df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "other_entity_id": [N_ROWS - 1 - i for i in range(N_ROWS)],
        }
    )
    feature_retrieval_job = client.get_batch_features(
        entity_rows=entity_df, feature_refs=["project1/feature_value6:1", "project1/other_feature_value7:1"]
    )
    output = feature_retrieval_job.to_dataframe()
    print(output.head())

    assert output["entity_id"].to_list() == [int(i) for i in output["project1_feature_value6_v1"].to_list()]
    assert output["other_entity_id"].to_list() == output["project1_other_feature_value7_v1"].to_list()

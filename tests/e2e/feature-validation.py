import pandas as pd
import pytest
import pytz
import uuid
import time
from datetime import datetime

from feast.client import Client
from feast.entity import Entity
from feast.feature import Feature
from feast.feature_set import FeatureSet
from feast.type_map import ValueType
from google.protobuf import json_format
from google.protobuf.duration_pb2 import Duration
from tensorflow_metadata.proto.v0 import statistics_pb2
import tensorflow_data_validation as tfdv


pd.set_option("display.max_columns", None)

PROJECT_NAME = "batch_" + uuid.uuid4().hex.upper()[0:6]


@pytest.fixture(scope="module")
def client(core_url, allow_dirty):
    # Get client for core and serving
    client = Client(core_url=core_url)
    client.create_project(PROJECT_NAME)
    client.set_project(PROJECT_NAME)

    # Ensure Feast core is active, but empty
    if not allow_dirty:
        feature_sets = client.list_feature_sets()
        if len(feature_sets) > 0:
            raise Exception(
                "Feast cannot have existing feature sets registered. Exiting tests."
            )

    return client


@pytest.fixture(scope="module")
def dataset_basic(client):
    fv_fs = FeatureSet(
        "feature_validation",
        features=[
            Feature("strings", ValueType.STRING),
            Feature("ints", ValueType.INT64),
            Feature("floats", ValueType.FLOAT),
        ],
        entities=[Entity("entity_id", ValueType.INT64)],
        max_age=Duration(seconds=100),
    )
    client.apply(fv_fs)
    time.sleep(20)

    N_ROWS = 20

    time_offset = datetime.utcnow().replace(tzinfo=pytz.utc)
    df = pd.DataFrame(
        {
            "datetime": [time_offset] * N_ROWS,
            "entity_id": [i for i in range(N_ROWS)],
            "strings": ["a", "b"] * int(N_ROWS / 2),
            "ints": [int(i) for i in range(N_ROWS)],
            "floats": [10.5 - i for i in range(N_ROWS)],
        }
    )

    expected_stats = tfdv.generate_statistics_from_dataframe(
        df[["entity_id", "strings", "ints", "floats"]]
    )
    clear_unsupported_fields(expected_stats)

    return {
        "id": client.ingest(fv_fs, df),
        "date": datetime.datetime(
            time_offset.year, time_offset.month, time_offset.day
        ).replace(tzinfo=pytz.utc),
        "stats": expected_stats,
    }


def test_basic_retrieval_by_single_dataset(client, dataset_basic):
    stats = client.get_statistics(
        features=["strings", "ints", "floats"],
        store="bigquery",
        dataset_ids=[dataset_basic["id"]],
    )

    assert stats == dataset_basic["stats"]


def test_basic_by_date(client, dataset_basic):
    stats = client.get_statistics(
        features=["strings", "ints", "floats"],
        store="bigquery",
        start_date=dataset_basic["date"],
        end_date=dataset_basic["date"],
    )
    assert stats == dataset_basic["stats"]


def clear_unsupported_fields(datasets):
    dataset = datasets.datasets[0]
    for feature in dataset.features:
        if feature.HasField("num_stats"):
            feature.num_stats.common_stats.ClearField("num_values_histogram")
        elif feature.HasField("string_stats"):
            feature.string_stats.common_stats.ClearField("num_values_histogram")
            for bucket in feature.string_stats.rank_histogram.buckets:
                bucket.ClearField("low_rank")
                bucket.ClearField("high_rank")
        elif feature.HasField("struct_stats"):
            feature.string_stats.struct_stats.ClearField("num_values_histogram")
        elif feature.HasField("bytes_stats"):
            feature.string_stats.bytes_stats.ClearField("num_values_histogram")

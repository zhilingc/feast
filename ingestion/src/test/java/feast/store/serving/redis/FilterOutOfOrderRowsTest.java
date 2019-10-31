package feast.store.serving.redis;

import static org.junit.Assert.*;

import com.google.protobuf.Timestamp;
import feast.core.FeatureSetProto.FeatureSetSpec;
import feast.store.serving.redis.RedisCustomIO.RedisMutation;
import feast.types.FeatureRowProto;
import feast.types.FeatureRowProto.FeatureRow;
import feast.types.FieldProto;
import feast.types.ValueProto;
import org.apache.beam.sdk.extensions.protobuf.ProtoCoder;
import org.apache.beam.sdk.testing.PAssert;
import org.apache.beam.sdk.testing.TestPipeline;
import org.apache.beam.sdk.testing.TestStream;
import org.apache.beam.sdk.values.PCollection;
import org.joda.time.Duration;
import org.joda.time.Instant;
import org.junit.Rule;
import org.junit.Test;

public class FilterOutOfOrderRowsTest {

  @Rule
  public transient TestPipeline p = TestPipeline.create();

  private FeatureRowProto.FeatureRow newFeatureRow(String featureSet, String fieldName, Integer value, Long secondsSinceEpoch) {
    return FeatureRowProto.FeatureRow.newBuilder()
        .setEventTimestamp(Timestamp.newBuilder().setSeconds(secondsSinceEpoch).build())
        .setFeatureSet(featureSet)
        .addFields(
            FieldProto.Field.newBuilder()
                .setName(fieldName)
                .setValue(ValueProto.Value.newBuilder().setInt32Val(value).build()))
        .build();
  }

  @Test
  public void shouldFilterOutdatedFeatureRow() {
    Duration expiryTime = Duration.standardSeconds(120);

    FeatureRowProto.FeatureRow feature1Recent = newFeatureRow("fs1", "fn", 1, 90L);
    FeatureRowProto.FeatureRow feature1Recent2 = newFeatureRow("fs1", "fn", 1, 95L);
    FeatureRowProto.FeatureRow feature1Outdated = newFeatureRow("fs1", "fn", 1, 80L);
    FeatureRowProto.FeatureRow feature1ResentAfterExpiry = newFeatureRow("fs1", "fn", 1, 85L);

    TestStream<FeatureRow> featureRowTestStream = TestStream.create(ProtoCoder.of(FeatureRowProto.FeatureRow.class))
        .advanceWatermarkTo(new Instant(0L))
        .addElements(feature1Recent)
        .addElements(feature1Recent2)
        .addElements(feature1Outdated)
//        .advanceWatermarkTo(new Instant(0L).plus(expiryTime.plus(1)))
        .addElements(feature1ResentAfterExpiry)
        .advanceWatermarkToInfinity();

    PCollection<RedisMutation> filtered = p.apply(featureRowTestStream)
        .apply(new FilterOutOfOrderRows(FeatureSetSpec.newBuilder().build()));
//    PAssert
//        .that(filtered).containsInAnyOrder(feature1Recent, feature2Recent, feature1ResentAfterExpiry);
    p.run();
  }
}
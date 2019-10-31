package feast.store.serving.redis;

import feast.core.FeatureSetProto.FeatureSetSpec;
import feast.storage.RedisProto.RedisKey;
import feast.store.serving.redis.RedisCustomIO.RedisMutation;
import feast.types.FeatureRowProto.FeatureRow;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.transforms.GroupByKey;
import org.apache.beam.sdk.transforms.PTransform;
import org.apache.beam.sdk.transforms.ParDo;
import org.apache.beam.sdk.transforms.windowing.AfterPane;
import org.apache.beam.sdk.transforms.windowing.AfterWatermark;
import org.apache.beam.sdk.transforms.windowing.FixedWindows;
import org.apache.beam.sdk.transforms.windowing.Repeatedly;
import org.apache.beam.sdk.transforms.windowing.Window;
import org.apache.beam.sdk.values.KV;
import org.apache.beam.sdk.values.PCollection;
import org.joda.time.Duration;
import org.joda.time.Instant;

public class FilterOutOfOrderRows extends
    PTransform<PCollection<FeatureRow>, PCollection<RedisMutation>> {

  private FeatureSetSpec featureSetSpec;

  public FilterOutOfOrderRows(FeatureSetSpec featureSetSpec) {
    this.featureSetSpec = featureSetSpec;
  }

  @Override
  public PCollection<RedisMutation> expand(PCollection<FeatureRow> input) {
    return input.apply(
        "FilterOutOfOrder",
        Window.<FeatureRow>into(FixedWindows.of(Duration.standardMinutes(1)))
            .triggering(Repeatedly.forever(AfterPane.elementCountAtLeast(1)))
            .withAllowedLateness(Duration.ZERO)
            .discardingFiredPanes()
    )
        .apply("MutateFeatureRows", ParDo.of(new FeatureRowToRedisMutationDoFn(featureSetSpec)))
        .apply("GroupByKey", GroupByKey.create())
        .apply("ExtractRows",
            ParDo.of(new DoFn<KV<RedisKey, Iterable<RedisMutation>>, RedisMutation>() {
              @ProcessElement
              public void processElement(ProcessContext c) {
                int ct = 0;
                RedisMutation latest = null;
                for (RedisMutation rm : c.element().getValue()) {
                  if (latest == null || rm.getEventTimestamp().getSeconds() > latest
                      .getEventTimestamp().getSeconds()) {
                    latest = rm;
                  }
                  ct ++;
                }
                System.out.println(String.format("pane %s: n = %s", c.pane().getIndex(), ct));
                if (latest != null) {
                  c.output(latest);
                }
              }
            }));
  }
}

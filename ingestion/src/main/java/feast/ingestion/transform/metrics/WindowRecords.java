/*
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2018-2019 The Feast Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package feast.ingestion.transform.metrics;

import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.transforms.GroupByKey;
import org.apache.beam.sdk.transforms.PTransform;
import org.apache.beam.sdk.transforms.ParDo;
import org.apache.beam.sdk.transforms.windowing.FixedWindows;
import org.apache.beam.sdk.transforms.windowing.Window;
import org.apache.beam.sdk.values.KV;
import org.apache.beam.sdk.values.PCollection;
import org.joda.time.Duration;

public class WindowRecords<T>
    extends PTransform<PCollection<T>, PCollection<KV<Integer, Iterable<T>>>> {

  private final long windowSize;

  public WindowRecords(long windowSize) {
    this.windowSize = windowSize;
  }

  @Override
  public PCollection<KV<Integer, Iterable<T>>> expand(PCollection<T> input) {
    return input
        .apply("Window records", Window.into(FixedWindows.of(Duration.standardSeconds(windowSize))))
        .apply(
            "Add key",
            ParDo.of(
                new DoFn<T, KV<Integer, T>>() {
                  @ProcessElement
                  public void processElement(ProcessContext c) {
                    c.output(KV.of(1, c.element()));
                  }
                }))
        .apply("Collect", GroupByKey.create());
  }
}

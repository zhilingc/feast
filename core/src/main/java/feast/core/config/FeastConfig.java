/*
 * Copyright 2018 The Feast Authors
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
 *
 */

package feast.core.config;

import java.util.HashMap;
import java.util.Map;
import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Getter
@Setter
@Component
@ConfigurationProperties(prefix = "feast")
public class FeastConfig {

  private JobConfig jobConfig;
  private StatsdConfig statsd;

  @Getter
  @Setter
  public static class JobConfig {

    private String coreApiUrl;
    private String runner;
    private String executable;
    private Map<String, String> options = new HashMap<>();
    private String errorsStoreType;
    private Map<String, String> errorsStoreOptions = new HashMap<>();
    private int monitoringPeriod;
    private int monitoringInitialDelay;
  }

  @Getter
  @Setter
  public static class StatsdConfig {

    private String host;
    private int port;
  }
}



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
package feast.core.service;

import com.google.protobuf.InvalidProtocolBufferException;
import feast.core.CoreServiceProto.ListFeatureSetsRequest;
import feast.core.CoreServiceProto.ListStoresRequest.Filter;
import feast.core.CoreServiceProto.ListStoresResponse;
import feast.core.FeatureSetProto.FeatureSetSpec;
import feast.core.StoreProto;
import feast.core.StoreProto.Store.Subscription;
import feast.core.dao.JobInfoRepository;
import feast.core.job.JobManager;
import feast.core.job.JobUpdateTask;
import feast.core.model.JobInfo;
import feast.core.model.JobStatus;
import feast.core.model.Source;
import feast.core.model.Store;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Optional;
import java.util.Set;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorCompletionService;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.stream.Collectors;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
public class JobCoordinatorService {

  private final long POLLING_INTERVAL_MILLISECONDS = 60000; // 1 min
  private JobInfoRepository jobInfoRepository;
  private SpecService specService;
  private JobManager jobManager;

  @Autowired
  public JobCoordinatorService(
      JobInfoRepository jobInfoRepository, SpecService specService, JobManager jobManager) {
    this.jobInfoRepository = jobInfoRepository;
    this.specService = specService;
    this.jobManager = jobManager;
  }

  /**
   * Poll does the following:
   *
   * <p>1) Checks DB and extracts jobs that have to run based on the specs available
   *
   * <p>2) Does a diff with the current set of jobs, starts/updates job(s) if necessary
   *
   * <p>3) Updates job object in DB with status, feature sets
   */
  @Transactional
  @Scheduled(fixedDelay = POLLING_INTERVAL_MILLISECONDS)
  public void Poll() {
    log.info("Polling for new jobs...");
    List<JobUpdateTask> jobUpdateTasks = new ArrayList<>();
    ListStoresResponse listStoresResponse = specService.listStores(Filter.newBuilder().build());
    for (StoreProto.Store store : listStoresResponse.getStoreList()) {
      Set<FeatureSetSpec> featureSetSpecs = new HashSet<>();
      try {
        for (Subscription subscription : store.getSubscriptionsList()) {
          featureSetSpecs.addAll(
              specService
                  .listFeatureSets(
                      ListFeatureSetsRequest.Filter.newBuilder()
                          .setFeatureSetName(subscription.getName())
                          .setFeatureSetVersion(subscription.getVersion())
                          .build())
                  .getFeatureSetsList());
        }
        if (!featureSetSpecs.isEmpty()) {
          featureSetSpecs.stream()
              .collect(Collectors.groupingBy(FeatureSetSpec::getSource))
              .entrySet()
              .stream()
              .forEach(
                  kv -> {
                    Optional<JobInfo> originalJob =
                        getJob(Source.fromProto(kv.getKey()), Store.fromProto(store));
                    jobUpdateTasks.add(
                        new JobUpdateTask(
                            kv.getValue(), kv.getKey(), store, originalJob, jobManager));
                  });
        }
      } catch (InvalidProtocolBufferException e) {
        log.warn("Unable to retrieve feature sets for store {}: {}", store, e.getMessage());
      }
    }
    if (jobUpdateTasks.size() == 0) {
      log.info("No jobs found.");
      return;
    }

    log.info("Creating/Updating {} jobs...", jobUpdateTasks.size());
    ExecutorService executorService = Executors.newFixedThreadPool(jobUpdateTasks.size());
    ExecutorCompletionService<JobInfo> ecs = new ExecutorCompletionService<>(executorService);
    jobUpdateTasks.forEach(ecs::submit);

    int completedTasks = 0;
    while (completedTasks < jobUpdateTasks.size()) {
      try {
        JobInfo jobInfo = ecs.take().get();
        if (jobInfo != null) {
          jobInfoRepository.saveAndFlush(jobInfo);
        }
      } catch (ExecutionException | InterruptedException e) {
        log.warn("Unable to start or update job: {}", e.getMessage());
      }
      completedTasks++;
    }
  }

  @Transactional
  public Optional<JobInfo> getJob(Source source, Store store) {
    List<JobInfo> jobs =
        jobInfoRepository.findBySourceIdAndStoreNameOrderByLastUpdatedDesc(
            source.getId(), store.getName());
    jobs =
        jobs.stream()
            .filter(job -> !JobStatus.getTerminalState().contains(job.getStatus()))
            .collect(Collectors.toList());
    if (jobs.size() == 0) {
      return Optional.empty();
    }
    // return the latest
    return Optional.of(jobs.get(0));
  }
}

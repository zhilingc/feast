"""
Main interface for users to interact with the Core API. 
"""

import feast.core.CoreService_pb2_grpc as core
from feast.sdk.resources.feature import Feature
from feast.sdk.resources.entity import Entity
from feast.sdk.resources.storage import Storage
from feast.sdk.resources.feature_group import FeatureGroup
import grpc


class Client:
    def __init__(self, serverURL, verbose=False):
        '''Create an instance of Feast client which is connected to feast 
        endpoint specified in the parameter
        
        Args:
            serverURI (string): feast's endpoint URL 
                                  (e.g.: "my.feast.com:8433")
        '''

        self.__channel = grpc.insecure_channel(serverURL)
        self.stub = core.CoreServiceStub(self.__channel)
        self.verbose = verbose

    def apply(self, obj):
        '''Create or update one or many feast's resource 
        (feature, entity, importer, storage).
        
        Args:
            object (object): one or many feast's resource
            create_entity (bool, optional):  (default: {None})
            create_features (bool, optional): [description] (default: {None})
        '''
        if isinstance(obj, list):
            ids = []
            for resource in obj:
                ids.append(self._apply(resource))
            return ids
        else:
            return self._apply(obj)

    def _apply(self, obj):
        '''Applies a single object to feast core.
        
        Args:
            obj (object): one of 
                           [Feature, Entity, FeatureGroup, Storage, Importer]
        '''
        if isinstance(obj, Feature):
            return self._apply_feature(obj)
        elif isinstance(obj, Entity):
            return self._apply_entity(obj)
        elif isinstance(obj, FeatureGroup):
            return self._apply_feature_group(obj)
        elif isinstance(obj, Storage):
            return self._apply_storage(obj)
        else:
            raise TypeError('Apply can only be passed one of the following \
            types: [Feature, Entity, FeatureGroup, Storage, Importer]')

    def _apply_feature(self, feature):
        '''Apply the feature to the core API
        
        Args:
            feature (Feature): feature to apply
        '''
        response = self.stub.ApplyFeature(feature.spec)
        if self.verbose: print("Successfully applied feature with id: {}\n---\n{}"
                .format(response.featureId, feature))
        return response.featureId

    def _apply_entity(self, entity):
        '''Apply the entity to the core API
        
        Args:
            entity (Entity): entity to apply
        '''
        response = self.stub.ApplyEntity(entity.spec)
        if self.verbose: print("Successfully applied entity with name: {}\n---\n{}"
            .format(response.entityName, entity))
        return response.entityName

    def _apply_feature_group(self, feature_group):
        '''Apply the feature group to the core API

        Args:
            feature_group (FeatureGroup): feature group to apply
        '''
        response = self.stub.ApplyFeatureGroup(feature_group.spec)
        if self.verbose: print("Successfully applied feature group with id: "+
            "{}\n---\n{}".format(response.featureGroupId, feature_group))
        return response.featureGroupId

    def _apply_storage(self, storage):
        '''Apply the storage to the core API
        
        Args:
            storage (Storage): storage to apply
        '''
        response = self.stub.ApplyStorage(storage.spec)
        if self.verbose: print("Successfully applied storage with id: "+
            "{}\n{}".format(response.storageId, storage))
        return response.storageId

    def close(self):
        self.channel.close()

    # def create_feature_set(self, entity=None, granularity=None, features=None):
    #     return feature_set()

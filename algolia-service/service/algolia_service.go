package service

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type AlgoliaService struct {
	ApplicationID string
	APIKey        string
	IndexName     string
	Client        algoliasearch.Client
	Index         algoliasearch.Index
}

func (self AlgoliaService) Init(applicationId string, APIkey string, indexName string) AlgoliaService {
	self.ApplicationID = applicationId
	self.APIKey = APIkey
	self.IndexName = indexName
	self.Client = algoliasearch.NewClient(applicationId, APIkey)
	self.Index = self.Client.InitIndex(indexName)
	return self
}

func (self AlgoliaService) Search(keyword string, mapParams algoliasearch.Map) (algoliasearch.QueryRes, error) {
	res, err := self.Index.Search(keyword, mapParams)
	return res, err
}

func (self AlgoliaService) AddObjects(objects []algoliasearch.Object) (algoliasearch.BatchRes, error) {
	res, err := self.Index.AddObjects(objects)
	return res, err
}

func (self AlgoliaService) UpdateObjects(objects []algoliasearch.Object) (algoliasearch.BatchRes, error) {
	res, err := self.Index.UpdateObjects(objects)
	return res, err
}

func (self AlgoliaService) DeleteObjects(objectIDs []string) (algoliasearch.BatchRes, error) {
	res, err := self.Index.DeleteObjects(objectIDs)
	return res, err
}

func (self AlgoliaService) GetObjects(objectIDs []string) ([]algoliasearch.Object, error) {
	res, err := self.Index.GetObjects(objectIDs)
	return res, err
}

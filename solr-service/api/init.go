package api

import (
	"github.com/autonomousdotai/handshake-services/solr-service/service"
	"github.com/autonomousdotai/handshake-services/solr-service/configs"
)

func CreateHandshakeSolrService() (service.SolrService) {
	sv := service.SolrService{}
	sv = sv.Init(configs.SolrHost, configs.SolrPort, configs.SolrCollectionHandshake)
	return sv
}

func CreateUserSolrService() (service.SolrService) {
	sv := service.SolrService{}
	sv = sv.Init(configs.SolrHost, configs.SolrPort, configs.SolrCollectionUser)
	return sv
}

var handshakeService = CreateHandshakeSolrService()
var userService = CreateUserSolrService()

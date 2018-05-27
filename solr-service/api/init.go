package api

import (
	"github.com/autonomousdotai/handshake-services/solr-service/service"
	"github.com/autonomousdotai/handshake-services/solr-service/setting"
)

func CreateHandshakeSolrService() (service.SolrService) {
	sv := service.SolrService{}
	sv = sv.Init(setting.CurrentConfig().SolrHost, setting.CurrentConfig().SolrPort, setting.CurrentConfig().SolrHandshakeCollection)
	return sv
}

func CreateUserSolrService() (service.SolrService) {
	sv := service.SolrService{}
	sv = sv.Init(setting.CurrentConfig().SolrHost, setting.CurrentConfig().SolrPort, setting.CurrentConfig().SolrUserCollection)
	return sv
}

var handshakeService = CreateHandshakeSolrService()
var userService = CreateUserSolrService()

package api

import (
	"github.com/autonomousdotai/handshake-services/algolia-service/service"
	"github.com/autonomousdotai/handshake-services/algolia-service/setting"
)

func CreateHandshakeAlgoliaService() (service.AlgoliaService) {
	sv := service.AlgoliaService{}
	sv = sv.Init(setting.CurrentConfig().AlgoliaApplicationID, setting.CurrentConfig().AlgoliaAPIKey, setting.CurrentConfig().AlgoliaHanshakeIndexName)
	return sv
}

func CreateUserAlgoliaService() (service.AlgoliaService) {
	sv := service.AlgoliaService{}
	sv = sv.Init(setting.CurrentConfig().AlgoliaApplicationID, setting.CurrentConfig().AlgoliaAPIKey, setting.CurrentConfig().AlgoliaUserIndexName)
	return sv
}

var handshakeService = CreateHandshakeAlgoliaService()
var userService = CreateUserAlgoliaService()

package api

import (
	"github.com/autonomousdotai/handshake-services/algolia-service/service"
	"github.com/autonomousdotai/handshake-services/algolia-service/setting"
)

func CreateAlgoliaService() (service.AlgoliaService) {
	sv := service.AlgoliaService{}
	sv = sv.Init(setting.CurrentConfig().AlgoliaApplicationID, setting.CurrentConfig().AlgoliaAPIKey, setting.CurrentConfig().AlgoliaIndexName)
	return sv
}

var algoliaService = CreateAlgoliaService()

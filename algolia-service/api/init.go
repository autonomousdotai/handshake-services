package api

import (
	"../service"
	"../setting"
)

// service

func CreateAlgoliaService() (service.AlgoliaService) {
	sv := service.AlgoliaService{}
	sv = algoliaService.Init(setting.CurrentConfig().AlgoliaApplicationID, setting.CurrentConfig().AlgoliaAPIKey, setting.CurrentConfig().AlgoliaIndexName)
	return sv
}

var algoliaService = service.AlgoliaService{}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type Api struct {
}

func (api Api) Init(router *gin.Engine) *gin.Engine {
	router.GET("/search", func(context *gin.Context) {
		api.Search(context)
	})
	router.GET("/objects", func(context *gin.Context) {
		api.GetObjects(context)
	})
	router.POST("/objects", func(context *gin.Context) {
		api.AddObjects(context)
	})
	router.PUT("/objects", func(context *gin.Context) {
		api.PartialUpdateObjects(context)
	})
	router.DELETE("/objects", func(context *gin.Context) {
		api.DeleteObjects(context)
	})
	return router
}

func (api Api) Search(context *gin.Context) {
	mapParams := new(algoliasearch.Map)
	err := context.Bind(&mapParams)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	keyword := context.Query("keyword")
	result, err := algoliaService.Search(keyword, *mapParams)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api Api) GetObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := algoliaService.GetObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api Api) AddObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := algoliaService.AddObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api Api) PartialUpdateObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := algoliaService.PartialUpdateObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api Api) DeleteObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := algoliaService.DeleteObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

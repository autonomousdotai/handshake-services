package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type HandshakeApi struct {
}

func (api HandshakeApi) Init(router *gin.Engine) *gin.RouterGroup {
	handshakeApi := router.Group("/handshake")
	{
		handshakeApi.GET("/search", func(context *gin.Context) {
			api.Search(context)
		})
		handshakeApi.GET("/objects", func(context *gin.Context) {
			api.GetObjects(context)
		})
		handshakeApi.POST("/objects", func(context *gin.Context) {
			api.AddObjects(context)
		})
		handshakeApi.PUT("/objects", func(context *gin.Context) {
			api.PartialUpdateObjects(context)
		})
		handshakeApi.DELETE("/objects", func(context *gin.Context) {
			api.DeleteObjects(context)
		})
	}
	return handshakeApi
}

func (api HandshakeApi) Search(context *gin.Context) {
	mapParams := new(algoliasearch.Map)
	err := context.Bind(&mapParams)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	keyword := context.Query("keyword")
	result, err := handshakeService.Search(keyword, *mapParams)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api HandshakeApi) GetObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.GetObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api HandshakeApi) AddObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.AddObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api HandshakeApi) PartialUpdateObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.PartialUpdateObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api HandshakeApi) DeleteObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.DeleteObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

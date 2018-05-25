package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type UserApi struct {
}

func (api UserApi) Init(router *gin.Engine) *gin.RouterGroup {
	handshakeApi := router.Group("/user")
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

func (api UserApi) Search(context *gin.Context) {
	mapParams := new(algoliasearch.Map)
	err := context.Bind(&mapParams)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	keyword := context.Query("keyword")
	result, err := userService.Search(keyword, *mapParams)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) GetObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.GetObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) AddObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.AddObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) PartialUpdateObjects(context *gin.Context) {
	objects := new([]algoliasearch.Object)
	err := context.Bind(&objects)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.PartialUpdateObjects(*objects)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) DeleteObjects(context *gin.Context) {
	objectIDs := new([]string)
	err := context.Bind(&objectIDs)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.DeleteObjects(*objectIDs)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

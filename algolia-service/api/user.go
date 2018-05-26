package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"log"
)

type UserApi struct {
}

func (api UserApi) Init(router *gin.Engine) *gin.RouterGroup {
	userApi := router.Group("/user")
	{
		userApi.GET("/search", func(context *gin.Context) {
			api.Search(context)
		})
		userApi.GET("/object", func(context *gin.Context) {
			api.GetObject(context)
		})
		userApi.GET("/objects", func(context *gin.Context) {
			api.GetObjects(context)
		})
		userApi.POST("/objects", func(context *gin.Context) {
			api.AddObjects(context)
		})
		userApi.PUT("/objects", func(context *gin.Context) {
			api.PartialUpdateObjects(context)
		})
		userApi.DELETE("/objects", func(context *gin.Context) {
			api.DeleteObjects(context)
		})
	}
	return userApi
}

func (api UserApi) Search(context *gin.Context) {
	mapParams := new(algoliasearch.Map)
	err := context.Bind(&mapParams)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	keyword := context.Query("keyword")
	result, err := userService.Search(keyword, *mapParams)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) GetObject(context *gin.Context) {
	objectID := context.Query("objectID")
	result, err := userService.GetObjects([]string{objectID})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, nil)
		return
	}
	if len(result) > 0 {
		context.JSON(http.StatusOK, result[0])
		return
	} else {
		context.JSON(http.StatusNotFound, nil)
		return
	}
}

func (api UserApi) GetObjects(context *gin.Context) {
	objectIDs := context.QueryArray("objectIDs")
	result, err := userService.GetObjects(objectIDs)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.AddObjects(*objects)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.PartialUpdateObjects(*objects)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.DeleteObjects(*objectIDs)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

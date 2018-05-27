package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rtt/Go-Solr"
)

type UserApi struct {
}

func (api UserApi) Init(router *gin.Engine) *gin.RouterGroup {
	userApi := router.Group("/user")
	{
		userApi.POST("/select", func(context *gin.Context) {
			api.Select(context)
		})
		userApi.POST("/update", func(context *gin.Context) {
			api.Update(context)
		})
	}
	return userApi
}

func (api UserApi) Select(context *gin.Context) {
	q := new(solr.Query)
	err := context.Bind(&q)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.Select(q)
	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api UserApi) Update(context *gin.Context) {
	document := new(map[string]interface{})
	err := context.Bind(&document)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := userService.Update(*document)
	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

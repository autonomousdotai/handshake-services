package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rtt/Go-Solr"
)

type HandshakeApi struct {
}

func (api HandshakeApi) Init(router *gin.Engine) *gin.RouterGroup {
	handshakeApi := router.Group("/handshake")
	{
		handshakeApi.POST("/select", func(context *gin.Context) {
			api.Select(context)
		})
		handshakeApi.POST("/update", func(context *gin.Context) {
			api.Update(context)
		})
	}
	return handshakeApi
}

func (api HandshakeApi) Select(context *gin.Context) {
	q := new(solr.Query)
	err := context.Bind(&q)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.Select(q)
	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

func (api HandshakeApi) Update(context *gin.Context) {
	document := new(map[string]interface{})
	err := context.Bind(&document)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := handshakeService.Update(*document)
	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

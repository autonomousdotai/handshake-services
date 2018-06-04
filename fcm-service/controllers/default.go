package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type DefaultController struct{}


func (d DefaultController) Home(c *gin.Context) {
    resp := JsonResponse{1, "IPFS REST API", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) Upload(c *gin.Context) {

    resp := JsonResponse{0, "Upload", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) View(c *gin.Context) {
    resp := JsonResponse{0, "View", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) NotFound(c *gin.Context) {
    resp := JsonResponse{0, "Page not found", nil}
    c.JSON(http.StatusOK, resp)
}

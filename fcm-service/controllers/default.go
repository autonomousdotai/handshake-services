package controllers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type DefaultController struct{}


func (d DefaultController) Home(c *gin.Context) {
    resp := JsonResponse{1, "IPFS REST API", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) Send(c *gin.Context) {
    var jsonData map[string]interface{}

    c.Bind(&jsonData)
    data, ok := jsonData["data"]
    
    if !ok {
        resp := JsonResponse{0, "Invalid data", nil}
        c.JSON(http.StatusOK, resp)
        c.Abort()
        return;
    }

    var status int
    var message string

    result, err := fcmService.Send(data.(map[string]interface{}))
 
    if result {
        status = 1
    } else {
        status = 0
        if err != nil {
            message = err.Error()
        }
    }

    resp := JsonResponse{status, message, nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) NotFound(c *gin.Context) {
    resp := JsonResponse{0, "Page not found", nil}
    c.JSON(http.StatusOK, resp)
}

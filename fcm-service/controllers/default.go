package controllers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gin-gonic/gin"
)

type DefaultController struct{}


func (d DefaultController) Home(c *gin.Context) {
    resp := JsonResponse{1, "IPFS REST API", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) Send(c *gin.Context) {
    var reqData map[string]interface{}

    c.Bind(&reqData)
    data, ok := reqData["data"]
    
    fmt.Println(reqData, data)

    if !ok {
        resp := JsonResponse{0, "Invalid data", nil}
        c.JSON(http.StatusOK, resp)
        c.Abort()
        return;
    }

    var status int
    var message string
    var jsonData map[string]interface{}

    json.Unmarshal([]byte(data.(string)), &jsonData)

    result, err := fcmService.Send(jsonData)
 
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

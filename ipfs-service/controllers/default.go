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
    file, fileErr := c.FormFile("file")

    if fileErr == nil {
        hash, err := ipfsService.Upload(file)
        var status int
        var message string
        
        if err != nil {
            status = 0
            message = err.Error()
        } else {
            status = 1
        }

        resp := JsonResponse{status, message, hash}
        c.JSON(http.StatusOK, resp)
        c.Abort()
        return;
    }

    resp := JsonResponse{0, "Invalid file", nil}
    c.JSON(http.StatusOK, resp)
}

func (d DefaultController) View(c *gin.Context) {
    hash := c.Param("hash")

    bytes, err := ipfsService.View(hash)

    if err != nil {
        resp := JsonResponse{0, err.Error(), nil}
        c.JSON(http.StatusOK, resp)
        c.Abort()
        return;
    }

    c.Data(http.StatusOK, "application/octet-stream", bytes)
}

func (d DefaultController) NotFound(c *gin.Context) {
    resp := JsonResponse{0, "Page not found", nil}
    c.JSON(http.StatusOK, resp)
}

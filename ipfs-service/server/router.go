package server

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "time"
    "log"
    "strconv"
    "github.com/gin-gonic/gin"

    "github.com/autonomousdotai/handshake-dispatcher/controllers"
    "github.com/autonomousdotai/handshake-dispatcher/config"
)

func NewRouter() *gin.Engine {
    router := gin.New()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    defaultController := new(controllers.DefaultController)
    router.GET("/", defaultController.Home) 

    router.NoRoute(defaultController.NotFound)

    return router
}

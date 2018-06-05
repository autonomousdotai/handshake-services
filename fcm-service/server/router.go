package server

import (
    "github.com/gin-gonic/gin"

    "github.com/ninjadotorg/handshake-services/fcm-service/controllers"
)

func NewRouter() *gin.Engine {
    router := gin.New()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    defaultController := new(controllers.DefaultController)
    router.GET("/", defaultController.Home) 
    router.POST("/send", defaultController.Send)
    router.NoRoute(defaultController.NotFound)

    return router
}

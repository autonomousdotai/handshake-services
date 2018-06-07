package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-services/comment-service/api"
	"github.com/ninjadotorg/handshake-services/comment-service/configs"
)

func main() {

	// Logger
	logFile, err := os.OpenFile("logs/autonomous_service.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter) // You may need this
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// end Logger
	// Setting router
	router := gin.Default()
	router.Use(Logger())
	router.Use(AuthorizeMiddleware())
	// Router Index
	index := router.Group("/")
	{
		index.GET("/", func(context *gin.Context) {
			result := map[string]interface{}{
				"status":  1,
				"message": "Comment Service API",
			}
			context.JSON(http.StatusOK, result)
		})
	}
	api := api.Api{}
	api.Init(router)
	router.Run(fmt.Sprintf(":%d", configs.ServicePort))
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		status := context.Writer.Status()
		latency := time.Since(t)
		log.Print("Request: " + context.Request.URL.String() + " | " + context.Request.Method + " - Status: " + strconv.Itoa(status) + " - " +
			latency.String())
	}
}

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, _ := strconv.ParseInt(context.GetHeader("Uid"), 10, 64)
		if userId <= 0 {
			context.JSON(http.StatusOK, gin.H{"status": 0, "message": "user is not logged in"})
			context.Abort()
			return
		}
		context.Set("UserId", userId)
		context.Next()
	}
}

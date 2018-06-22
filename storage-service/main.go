package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-services/storage-service/configs"
	gocontext "golang.org/x/net/context"
	"google.golang.org/api/option"
)

var gsBucket *storage.BucketHandle

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
	// Router Index

	if gsBucket == nil {
		opt := option.WithCredentialsFile(configs.GSCredentialsFile)
		ctx := gocontext.Background()
		client, err := storage.NewClient(ctx, opt)
		if err != nil {
			panic(err)
		}
		bucketName := configs.GSBucketName
		gsBucket = client.Bucket(bucketName)
	}

	index := router.Group("/")
	{
		index.GET("/", func(context *gin.Context) {
			result := map[string]interface{}{
				"status":  1,
				"message": "Storage Service API",
			}
			context.JSON(http.StatusOK, result)
		})
		index.POST("/", func(context *gin.Context) {
			file := context.Query("file")
			if file == "" {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "file is invalid",
				})
			}
			buffer, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				if err != nil {
					log.Print(err)
					context.JSON(http.StatusOK, gin.H{
						"status":  -1,
						"message": err.Error(),
					})
				}
			}

			if err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
			}
			fileBytes := bytes.NewReader(buffer)
			fileType := http.DetectContentType(buffer)
			ctx := gocontext.Background()
			w := gsBucket.Object(file).NewWriter(ctx)
			w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
			w.CacheControl = "public, max-age=86400"
			w.ContentType = fileType
			if _, err := io.Copy(w, fileBytes); err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
			}
			if err := w.Close(); err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  1,
				"message": "OK",
			})
		})

		index.POST("/user/upload", func(context *gin.Context) {
			userID, _ := strconv.ParseInt(context.GetHeader("Uid"), 10, 64)
			if userID <= 0 {
				context.JSON(http.StatusOK, gin.H{"status": 0, "message": "User is not authorized"})
				return
			}

			file := context.Query("file")
			if file == "" {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "file is invalid",
				})
				return
			}
			file = fmt.Sprintf("users/%d/%s", userID, file)
			buffer, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				if err != nil {
					log.Print(err)
					context.JSON(http.StatusOK, gin.H{
						"status":  -1,
						"message": err.Error(),
					})
					return
				}
			}

			if err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				return
			}
			fileBytes := bytes.NewReader(buffer)
			fileType := http.DetectContentType(buffer)
			ctx := gocontext.Background()
			w := gsBucket.Object(file).NewWriter(ctx)
			w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
			w.CacheControl = "public, max-age=86400"
			w.ContentType = fileType
			if _, err := io.Copy(w, fileBytes); err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				return
			}
			if err := w.Close(); err != nil {
				log.Print(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  1,
				"message": "OK",
				"data":    file,
			})
			return
		})

	}
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

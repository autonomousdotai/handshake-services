package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/autonomousdotai/handshake-services/mail-service/configs"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/storage"
	"io/ioutil"
	"encoding/base64"
	"path/filepath"
	"strings"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"errors"
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

	index := router.Group("/")
	{
		index.GET("/", func(context *gin.Context) {
			result := map[string]interface{}{
				"status":  1,
				"message": "Mail Service API",
			}
			context.JSON(http.StatusOK, result)
		})
		index.POST("/", func(context *gin.Context) {
			from := context.Request.PostFormValue("from")
			to := context.Request.PostFormValue("to")
			subject := context.Request.PostFormValue("subject")
			body := context.Request.PostFormValue("body")
			err = Send(from, to, subject, body, context)
			if err != nil {
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
	}
	router.Run(fmt.Sprintf(":%d", configs.ServicePort))
}

func parseEmailAddress(email string) (name, address string) {
	strs := strings.Split(email, "<")
	if (len(strs) > 1) {
		address = strs[1]
		address = strings.Replace(address, ">", "", 1000)
		address = strings.TrimSpace(address)
		name = strs[0]
		name = strings.TrimSpace(name)
		return name, address
	} else {
		strs = strings.Split(email, "-")
		if (len(strs) > 1) {
			address = strs[1]
			address = strings.TrimSpace(address)
			name = strs[0]
			name = strings.TrimSpace(name)
			return name, address
		} else {
			address = email
			name = email
			return name, address
		}
	}
}

func Send(from, to, subject, body string, context *gin.Context) (error) {
	m := mail.NewV3Mail()
	name, address := parseEmailAddress(from)
	e := mail.NewEmail(name, address)
	m.SetFrom(e)
	m.Subject = subject
	p := mail.NewPersonalization()
	tos := []*mail.Email{mail.NewEmail("", to)}
	p.AddTos(tos...)
	m.AddPersonalizations(p)
	c := mail.NewContent("text/html", body)
	m.AddContent(c)
	form, _ := context.MultipartForm()
	files := form.File["file[]"]
	for _, file := range files {
		fileName := file.Filename
		sourceFile, err := file.Open()
		a := mail.NewAttachment()
		buffer, err := ioutil.ReadAll(sourceFile)
		if err != nil {
			fmt.Println("Read file error: ", err)
			return err
		}
		if err != nil {
			log.Println(err)
		} else {
			a.SetContent(base64.StdEncoding.EncodeToString(buffer))
			a.SetType("application/" + filepath.Ext(fileName))
			a.SetFilename(filepath.Base(fileName))
			a.SetDisposition("attachment")
			m.AddAttachment(a)
		}
	}
	Body := mail.GetRequestBody(m)
	request := sendgrid.GetRequest(configs.SendgridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return err
	} else {
		if response.StatusCode != 202 {
			return errors.New(string(response.Body))
		}
		return nil
	}
	return nil
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

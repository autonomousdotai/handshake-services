package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/ninjadotorg/handshake-services/mail-service/configs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/base64"
	"path/filepath"
	"strings"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
			form, err := context.MultipartForm()
			if err != nil {
				log.Println(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				return
			}

			from := context.Request.PostFormValue("from")
			if len(from) <= 0 {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "from is invalid",
				})
				return
			}
			to, ok := form.Value["to[]"]
			if !ok {
				to = []string{}
			}
			if len(to) <= 0 {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "to is invalid",
				})
				return
			}
			cc, ok := form.Value["cc[]"]
			if !ok {
				cc = []string{}
			}
			bcc, ok := form.Value["bcc[]"]
			if !ok {
				bcc = []string{}
			}
			subject := context.Request.PostFormValue("subject")
			if len(subject) <= 0 {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "subject is invalid",
				})
				return
			}
			body := context.Request.PostFormValue("body")
			if len(body) <= 0 {
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": "body is invalid",
				})
				return
			}
			files := form.File["file[]"]

			m := mail.NewV3Mail()
			name, address := parseEmailAddress(from)
			e := mail.NewEmail(name, address)
			m.SetFrom(e)
			m.Subject = subject
			p := mail.NewPersonalization()

			tos := make([]*mail.Email, 0)
			for _, _to := range to {
				tos = append(tos, mail.NewEmail("", _to))
			}
			p.AddTos(tos...)

			ccs := make([]*mail.Email, 0)
			for _, _cc := range cc {
				ccs = append(ccs, mail.NewEmail("", _cc))
			}
			p.AddCCs(ccs...)

			bccs := make([]*mail.Email, 0)
			for _, _bcc := range bcc {
				bccs = append(bccs, mail.NewEmail("", _bcc))
			}
			p.AddBCCs(bccs...)

			m.AddPersonalizations(p)
			c := mail.NewContent("text/html", body)
			m.AddContent(c)

			for _, file := range files {
				fileName := file.Filename
				sourceFile, err := file.Open()
				a := mail.NewAttachment()
				buffer, err := ioutil.ReadAll(sourceFile)
				if err != nil {
					log.Println(err)
					context.JSON(http.StatusOK, gin.H{
						"status":  -1,
						"message": err.Error(),
					})
					return
				}
				a.SetContent(base64.StdEncoding.EncodeToString(buffer))
				a.SetType("application/" + filepath.Ext(fileName))
				a.SetFilename(filepath.Base(fileName))
				a.SetDisposition("attachment")
				m.AddAttachment(a)
			}
			Body := mail.GetRequestBody(m)
			request := sendgrid.GetRequest(configs.SendgridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
			request.Method = "POST"
			request.Body = Body
			response, err := sendgrid.API(request)
			if err != nil {
				log.Println(err)
				context.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				return
			} else {
				if response.StatusCode != 202 {
					log.Print(err)
					context.JSON(http.StatusOK, gin.H{
						"status":  -1,
						"message": string(response.Body),
					})
					return
				}
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

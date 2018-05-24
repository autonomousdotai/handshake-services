package api

import (
	"github.com/gin-gonic/gin"
	"../response_obj"
	"../request_obj"
	"net/http"
	"../bean"
	"log"
)

type Api struct {
}

func (api Api) Init(router *gin.Engine) *gin.RouterGroup {
	apiGroupApi := router.Group("/api")
	{
		apiGroupApi.GET("/", func(context *gin.Context) {
			context.String(200, "Comment API")
		})
		apiGroupApi.POST("/comment", func(context *gin.Context) {
			api.CreateComment(context)
		})
	}
	return apiGroupApi
}

func (api Api) CreateComment(context *gin.Context) {
	result := new(response_obj.ResponseObject)

	userId, ok := context.Get("UserId")
	if !ok {
		result.SetStatus(bean.NotSignIn)
		context.JSON(http.StatusOK, result)
		return
	}
	if userId.(int64) <= 0 {
		result.SetStatus(bean.NotSignIn)
		context.JSON(http.StatusOK, result)
		return
	}

	request := new(request_obj.CommentRequest)
	err := context.Bind(&request)
	if err != nil {
		log.Print(err)
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}
	comment, appErr := commentService.CreateComment(userId.(int64), *request)
	if appErr != nil {
		log.Print(appErr.OrgError)
		result.SetStatus(bean.UnexpectedError)
		result.Error = appErr.OrgError.Error()
		context.JSON(http.StatusOK, result)
		return
	}
	data := response_obj.MakeCommentResponse(comment)

	result.Data = data
	result.Status = 1
	result.Message = ""
	context.JSON(http.StatusOK, result)
	return
}

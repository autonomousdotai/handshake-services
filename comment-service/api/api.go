package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-services/comment-service/response_obj"
	"github.com/ninjadotorg/handshake-services/comment-service/request_obj"
	"net/http"
	"github.com/ninjadotorg/handshake-services/comment-service/bean"
	"log"
	"strconv"
	"github.com/ninjadotorg/handshake-services/comment-service/utils"
	"encoding/json"
)

type Api struct {
}

func (api Api) Init(router *gin.Engine) *gin.Engine {
	router.GET("/list", func(context *gin.Context) {
		api.GetComments(context)
	})
	router.POST("/", func(context *gin.Context) {
		api.CreateComment(context)
	})
	router.GET("/count", func(context *gin.Context) {
		api.GetCommentCount(context)
	})
	return router
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

	requestJson := context.Request.PostFormValue("request")
	err := json.Unmarshal([]byte(requestJson), &request)
	if err != nil {
		log.Print(err)
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}
	sourceFile, sourceFileHeader, err := context.Request.FormFile("image")
	comment, appErr := commentService.CreateComment(userId.(int64), *request, &sourceFile, sourceFileHeader)
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

func (api Api) GetComments(context *gin.Context) {
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

	pageSizeStr := context.Query("page_size")
	if len(pageSizeStr) == 0 {
		pageSizeStr = utils.DEFAULT_PAGE_SIZE
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Print(err)
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}
	pageStr := context.Query("page")
	if len(pageStr) == 0 {
		pageStr = utils.DEFAULT_PAGE
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Print(err)
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}
	objectType := context.Query("object_type")
	objectId := context.Query("object_id")

	var pagination *bean.Pagination
	pagination = &bean.Pagination{PageSize: pageSize, Page: page}

	pagination, err = commentService.GetCommentPagination(0, objectType, objectId, pagination)
	if err != nil {
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}

	data := response_obj.MakePaginationCommentResponse(pagination)

	result.Data = data
	result.Status = 1
	result.Message = ""
	context.JSON(http.StatusOK, result)
	return
}

func (api Api) GetCommentCount(context *gin.Context) {
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

	objectType := context.Query("object_type")
	objectId, _ := strconv.ParseInt(context.Query("object_id"), 10, 64)
	userIdR, _ := strconv.ParseInt(context.Query("user_id"), 10, 64)

	count, err := commentService.GetCommentCount(objectType, objectId, userIdR)
	if err != nil {
		result.SetStatus(bean.UnexpectedError)
		result.Error = err.Error()
		context.JSON(http.StatusOK, result)
		return
	}

	result.Data = count
	result.Status = 1
	result.Message = ""
	context.JSON(http.StatusOK, result)
	return
}

package service

import (
	"github.com/ninjadotorg/handshake-services/comment-service/models"
	"github.com/ninjadotorg/handshake-services/comment-service/bean"
	"errors"
	"log"
	"github.com/ninjadotorg/handshake-services/comment-service/request_obj"
	"github.com/ninjadotorg/handshake-services/comment-service/configs"
	"mime/multipart"
	"strings"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

type CommentService struct {
}

func (commentService CommentService) CreateComment(userId int64, request request_obj.CommentRequest, sourceFile *multipart.File, sourceFileHeader *multipart.FileHeader) (models.Comment, *bean.AppError) {
	tx := models.Database().Begin()

	comment := models.Comment{}

	comment.UserId = userId
	comment.ObjectType = request.ObjectType
	comment.ObjectId = request.ObjectId
	comment.Comment = request.Comment
	comment.Status = 1

	comment, err := commentDao.Create(comment, tx)
	if err != nil {
		log.Println(err)

		tx.Rollback()
		return comment, &bean.AppError{errors.New(err.Error()), "Error occurred, please try again", -1, "error_occurred"}
	}

	filePath := ""
	if sourceFile != nil && sourceFileHeader != nil {
		//upload image for comment
		uploadImageFolder := "comment"
		fileName := sourceFileHeader.Filename
		imageExt := strings.Split(fileName, ".")[1]
		fileNameImage := fmt.Sprintf("comment-%d-image-%s.%s", comment.ID, time.Now().Format("20060102150405"), imageExt)
		filePath = uploadImageFolder + "/" + fileNameImage
		err := fileUploadService.UploadFile(filePath, sourceFile)
		if err != nil {
			log.Println(err)

			tx.Rollback()
			return comment, &bean.AppError{errors.New(err.Error()), "Error occurred, please try again", -1, "error_occurred"}
		}
	}

	comment.Image = filePath
	comment, err = commentDao.Update(comment, tx)
	if err != nil {
		log.Println(err)
		return comment, &bean.AppError{errors.New(err.Error()), "Error occurred, please try again", -1, "error_occurred"}
	}

	tx.Commit()

	return comment, nil
}

func (commentService CommentService) GetCommentPagination(userId int64, objectType string, objectId int64, pagination *bean.Pagination) (*bean.Pagination, error) {
	pagination, err := commentDao.GetCommentPagination(userId, objectType, objectId, pagination)
	comments := pagination.Items.([]models.Comment)
	items := []models.Comment{}
	for _, comment := range comments {
		user, _ := commentService.GetUser(comment.UserId)
		comment.User = user
		items = append(items, comment)
	}
	pagination.Items = items
	return pagination, err
}

func (commentService CommentService) GetUser(userId int64) (models.User, error) {
	result := models.JsonUserResponse{}
	url := fmt.Sprintf("%s/%d", configs.DispatcherServiceUrl+"/system/user", userId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return result.Data, err
	}
	req.Header.Set("Content-Type", "application/json")
	bodyBytes, err := netUtil.CurlRequest(req)
	if err != nil {
		log.Println(err)
		return result.Data, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Println(err)
		return result.Data, err
	}
	return result.Data, err
}

func (commentService CommentService) GetCommentCount(objectType string, objectId int64, userId int64) (int, error) {
	return commentDao.GetCommentCount(objectType, objectId, userId)
}

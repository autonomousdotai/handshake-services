package service

import (
	"github.com/autonomousdotai/handshake-services/comment-service/models"
	"github.com/autonomousdotai/handshake-services/comment-service/bean"
	"errors"
	"log"
	"github.com/autonomousdotai/handshake-services/comment-service/request_obj"
)

type CommentService struct {
}

func (commentService CommentService) CreateComment(userId int64, request request_obj.CommentRequest) (models.Comment, *bean.AppError) {
	crowdFunding := models.Comment{}

	crowdFunding.UserId = userId
	crowdFunding.ObjectType = request.ObjectType
	crowdFunding.ObjectId = request.ObjectId
	crowdFunding.Comment = request.Comment
	crowdFunding.Status = 1

	crowdFunding, err := commentDao.Create(crowdFunding, nil)
	if err != nil {
		log.Println(err)
		return crowdFunding, &bean.AppError{errors.New(err.Error()), "Error occurred, please try again", -1, "error_occurred"}
	}

	return crowdFunding, nil
}

func (commentService CommentService) GetCommentPagination(userId int64, objectType string, objectId int64, pagination *bean.Pagination) (*bean.Pagination, error) {
	pagination, err := commentDao.GetCommentPagination(userId, objectType, objectId, pagination)
	return pagination, err
}

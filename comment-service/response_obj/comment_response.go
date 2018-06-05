package response_obj

import (
	"time"
	"github.com/ninjadotorg/handshake-services/comment-service/models"
	"github.com/ninjadotorg/handshake-services/comment-service/utils"
	"github.com/ninjadotorg/handshake-services/comment-service/bean"
)

type UserResponse struct {
	ID     int64  `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Status int    `json:"status"`
}

type CommentResponse struct {
	DateCreated time.Time    `json:"date_created"`
	ID          int64        `json:"id"`
	UserId      int64        `json:"user_id"`
	ObjectId    string       `json:"object_id"`
	Comment     string       `json:"comment"`
	Image       string       `json:"image"`
	Address     string       `json:"address"`
	Status      int          `json:"status"`
	User        UserResponse `json:"user"`
}

func MakeCommentResponse(model models.Comment) CommentResponse {
	result := CommentResponse{}
	result.ID = model.ID
	result.UserId = model.UserId
	result.ObjectId = model.ObjectId
	result.Comment = model.Comment
	result.Image = utils.CdnUrlFor(model.Image)
	result.Status = model.Status
	result.DateCreated = model.DateCreated
	result.User = MakeUserResponse(model.User)
	return result
}

func MakeUserResponse(model models.User) UserResponse {
	result := UserResponse{}
	result.ID = model.ID
	result.Email = model.Email
	result.Name = model.Name
	result.Avatar = utils.CdnUrlFor(model.Avatar)
	return result
}

func MakeArrayCommentResponse(models []models.Comment) []CommentResponse {
	results := []CommentResponse{}
	for _, model := range models {
		result := MakeCommentResponse(model)
		results = append(results, result)
	}
	return results
}

func MakePaginationCommentResponse(pagination *bean.Pagination) PaginationResponse {
	return MakePaginationResponse(pagination.Page, pagination.PageSize, pagination.Total, MakeArrayCommentResponse(pagination.Items.([]models.Comment)))
}

type CountResponse struct {
	ObjectType string `json:"object_type"`
	ObjectId   int64  `json:"object_id"`
	UserId     int64  `json:"user_id"`
	User       int    `json:"count"`
}

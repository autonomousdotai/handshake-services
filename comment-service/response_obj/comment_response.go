package response_obj

import (
	"time"
	"github.com/autonomousdotai/handshake-services/comment-service/models"
	"github.com/autonomousdotai/handshake-services/comment-service/utils"
	"github.com/autonomousdotai/handshake-services/comment-service/bean"
)

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CommentResponse struct {
	DateCreated time.Time `json:"date_created"`
	ID          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Comment     string    `json:"comment"`
	Image       string    `json:"image"`
	Status      int       `json:"status"`
}

func MakeCommentResponse(model models.Comment) CommentResponse {
	result := CommentResponse{}
	result.ID = model.ID
	result.UserId = model.UserId
	result.Comment = model.Comment
	result.Image = utils.CdnUrlFor(model.Image)
	result.Status = model.Status
	result.DateCreated = model.DateCreated
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

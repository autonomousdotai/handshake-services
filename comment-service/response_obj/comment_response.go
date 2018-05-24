package response_obj

import (
	"time"
	"../models"
	"../utils"
)

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

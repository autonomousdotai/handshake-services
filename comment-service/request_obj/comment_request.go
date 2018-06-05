package request_obj

type CommentRequest struct {
	Comment  string `json:"comment"`
	ObjectId string `json:"object_id"`
	Address  string `json:"address"`
}

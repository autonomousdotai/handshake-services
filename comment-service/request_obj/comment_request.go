package request_obj

type CommentRequest struct {
	Comment    string `json:"comment"`
	ObjectType string `json:"object_type"`
	ObjectId   string `json:"object_id"`
}

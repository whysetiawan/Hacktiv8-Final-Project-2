package dto

type Comment struct {
	ID        uint   `json:"id"`
	CommentId uint   `json:"comment_id"`
	PhotoId   uint   `json:"photo_id"`
	Message   string `json:"message"`
}

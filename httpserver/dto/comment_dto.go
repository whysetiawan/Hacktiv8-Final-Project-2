package dto

type CreateCommentDto struct {
	PhotoId uint   `json:"photo_id"`
	Message string `json:"message"`
}

type UpdateCommentDto struct {
	Message string `json:"message"`
}

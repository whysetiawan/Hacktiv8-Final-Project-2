package models

type CommentModel struct {
	BaseModel
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

func (CommentModel) TableName() string {
	return "public.Comment"
}

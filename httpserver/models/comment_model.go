package models

type CommentModel struct {
	BaseModel
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

func (CommentModel) TableName() string {
	return "public.Comment"
}

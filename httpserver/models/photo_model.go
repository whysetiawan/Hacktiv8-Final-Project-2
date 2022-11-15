package models

type PhotoModel struct {
	BaseModel
	Title        string         `json:"title"`
	PhotoUrl     string         `json:"photo_url"`
	UserID       uint           `json:"user_id"`
	CommentModel []CommentModel `json:"comment" gorm:"foreignKey:PhotoID;references:ID"`
}

func (PhotoModel) TableName() string {
	return "public.Photo"
}

package models

import (
	"gorm.io/gorm"
)

type PhotoModel struct {
	gorm.Model
	PhotoId      uint           `gorm:"primaryKey;column:id" json:"photo_id"`
	Title        string         `json:"title"`
	PhotoUrl     string         `json:"photo_url"`
	UserId       uint           `json:"user_id"`
	CommentModel []CommentModel `json:"comment" gorm:"foreignKey:CommentId"`
}

func (PhotoModel) TableName() string {
	return "public.Photo"
}

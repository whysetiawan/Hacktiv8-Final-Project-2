package models

import (
	"gorm.io/gorm"
)

type CommentModel struct {
	gorm.Model
	CommentId uint   `gorm:"primaryKey;column:id" json:"photo_id"`
	Message   string `json:"message"`
	UserId    uint   `json:"user_id"`
}

func (CommentModel) TableName() string {
	return "public.Comment"
}

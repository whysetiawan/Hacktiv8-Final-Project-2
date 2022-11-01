package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserId       uint               `gorm:"primaryKey;column:id"  json:"user_id"`
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	Age          uint8              `json:"age"`
	SocialMedia  []SocialMediaModel `json:"social_medias" gorm:"foreignKey:SocialMediaId"`
	PhotoModel   []PhotoModel       `json:"photos" gorm:"foreignKey:PhotoId"`
	CommentModel []CommentModel     `json:"comments" gorm:"foreignKey:CommentId"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type UserParams struct {
	ID int64 `uri:"id" binding:"required"`
}

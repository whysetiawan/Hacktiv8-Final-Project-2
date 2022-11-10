package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username     string             `gorm:"uniqueIndex" json:"username"`
	Email        string             `gorm:"uniqueIndex" json:"email"`
	Password     string             `json:"password"`
	Age          uint8              `json:"age"`
	SocialMedia  []SocialMediaModel `json:"social_medias" gorm:"foreignKey:SocialMediaId"`
	PhotoModel   []PhotoModel       `json:"photos" gorm:"foreignKey:PhotoId"`
	CommentModel []CommentModel     `json:"comments" gorm:"foreignKey:CommentId"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

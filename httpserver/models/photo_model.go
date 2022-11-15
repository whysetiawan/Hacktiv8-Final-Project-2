package models

import (
	"time"

	"gorm.io/gorm"
)

type PhotoModel struct {
	gorm.Model
	PhotoId      uint           `gorm:"primaryKey;column:id" json:"photo_id"`
	Title        string         `json:"title" binding:"required"`
	Caption      string         `json:"caption" binding:"required"`
	PhotoUrl     string         `json:"photo_url" binding:"required"`
	UserId       uint           `json:"user_id"`
	CommentModel []CommentModel `json:"comment,omitempty" gorm:"foreignKey:CommentId"`
}

// type Photo struct {
// 	ID        uint   `gorm:"primayKey;column:id" json:"id"`
// 	Title     string `json:"title" binding:"required"`
// 	Caption   string `json:"caption" binding:"required"`
// 	PhotoUrl  string `json:"photoUrl" binding:"required"`
// 	UserId    int64  `json:"userId"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type PhotoModelResponse struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Caption   string    `json:"caption,omitempty"`
	PhotoUrl  string    `json:"photo_url,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"date,omitempty"`
}

func (PhotoModel) TableName() string {
	return "public.Photo"
}

type GetPhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"User" gorm:"foreignKey:id"`
}

type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
}

type UpdatePhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserID    int64     `json:"userID"`
	UpdatedAt time.Time `json:"updatedAt"`
}

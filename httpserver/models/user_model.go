package models

type UserModel struct {
	BaseModel
	Username     string             `gorm:"uniqueIndex" json:"username"`
	Email        string             `gorm:"uniqueIndex" json:"email"`
	Password     string             `json:"password"`
	Age          uint8              `json:"age"`
<<<<<<< HEAD
	SocialMedia  []SocialMediaModel `json:"social_medias" gorm:"foreignKey:UserID;references:ID"`
	PhotoModel   []PhotoModel       `json:"photos" gorm:"foreignKey:UserID;references:ID"`
	CommentModel []CommentModel     `json:"comments" gorm:"foreignKey:UserID;references:ID"`
=======
	SocialMedia  []SocialMediaModel `json:"social_medias" gorm:"foreignKey:UserId;references:ID"`
	PhotoModel   []PhotoModel       `json:"photos" gorm:"foreignKey:UserId;references:ID"`
	CommentModel []CommentModel     `json:"comments" gorm:"foreignKey:UserId;references:ID"`
>>>>>>> 9608a17a1a02f56d3ed18d6aa91a31f0730a7a73
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

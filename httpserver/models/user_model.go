package models

type UserModel struct {
	BaseModel
	Username     string             `gorm:"uniqueIndex" json:"username"`
	Email        string             `gorm:"uniqueIndex" json:"email"`
	Password     string             `json:"password"`
	Age          uint8              `json:"age"`
	SocialMedia  []SocialMediaModel `json:"social_medias" gorm:"foreignKey:UserID;references:ID"`
	PhotoModel   []PhotoModel       `json:"photos" gorm:"foreignKey:UserID;references:ID"`
	CommentModel []CommentModel     `json:"comments" gorm:"foreignKey:UserID;references:ID"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

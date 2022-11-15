package models

type SocialMediaModel struct {
	BaseModel
	SocialMediaId  uint   `gorm:"primaryKey;column:id" json:"social_media_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

func (SocialMediaModel) TableName() string {
	return "public.SocialMedia"
}

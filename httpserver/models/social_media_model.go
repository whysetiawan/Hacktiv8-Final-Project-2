package models

type SocialMediaModel struct {
	BaseModel
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

func (SocialMediaModel) TableName() string {
	return "public.SocialMedia"
}

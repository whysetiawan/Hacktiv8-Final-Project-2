package dto

type UpsertSocialMediaDto struct {
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
	Name           string `json:"name" binding:"required"`
}

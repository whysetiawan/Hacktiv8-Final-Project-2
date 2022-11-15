package repositories

import (
	"final-project-2/httpserver/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) CreateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error) {
	err := r.db.Create(socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

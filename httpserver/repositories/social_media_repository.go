package repositories

import (
	"final-project-2/httpserver/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error)
	GetSocialMedias(userID uint) (*[]models.SocialMediaModel, error)
	UpdateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error)
	DeleteSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) CreateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error) {
	err := r.db.Create(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *socialMediaRepository) GetSocialMedias(userID uint) (*[]models.SocialMediaModel, error) {
	var socialMedias []models.SocialMediaModel
	err := r.db.Find(&socialMedias).Order("id desc").Error

	if err != nil {
		return &socialMedias, err
	}

	return &socialMedias, nil
}

func (r *socialMediaRepository) UpdateSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error) {
	err := r.db.Model(socialMedia).Updates(socialMedia).Error

	if err != nil {
		return socialMedia, err
	}
	return socialMedia, err
}

func (r *socialMediaRepository) DeleteSocialMedia(socialMedia *models.SocialMediaModel) (*models.SocialMediaModel, error) {
	err := r.db.Model(socialMedia).Delete(socialMedia).Error

	if err != nil {
		return socialMedia, err
	}
	return socialMedia, err
}

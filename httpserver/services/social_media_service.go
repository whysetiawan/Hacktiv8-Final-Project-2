package services

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"
	"fmt"
)

type SocialMediaService interface {
	CreateSocialMedia(dto *dto.UpsertSocialMediaDto, userID uint) (*models.SocialMediaModel, error)
	GetSocialMedias(userID uint) (*[]models.SocialMediaModel, error)
	UpdateSocialMedia(dto *dto.UpsertSocialMediaDto, socialMediaID uint, userID uint) (*models.SocialMediaModel, error)
	DeleteSocialMedia(userID uint, socialMediaID uint) (*models.SocialMediaModel, error)
}

type socialMediaService struct {
	socialMediaRepository repositories.SocialMediaRepository
}

func NewSocialMediaService(r repositories.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{r}
}

func (s *socialMediaService) CreateSocialMedia(dto *dto.UpsertSocialMediaDto, userID uint) (*models.SocialMediaModel, error) {
	socialMedia := models.SocialMediaModel{
		UserID:         userID,
		Name:           dto.Name,
		SocialMediaUrl: dto.SocialMediaUrl,
	}
	fmt.Println(socialMedia)

	result, err := s.socialMediaRepository.CreateSocialMedia(&socialMedia)

	if err != nil {
		return &socialMedia, err
	}

	return result, nil

}

func (s *socialMediaService) GetSocialMedias(userID uint) (*[]models.SocialMediaModel, error) {

	result, err := s.socialMediaRepository.GetSocialMedias(userID)

	if err != nil {
		return result, err
	}

	return result, nil

}

func (s *socialMediaService) UpdateSocialMedia(dto *dto.UpsertSocialMediaDto, socialMediaID uint, userID uint) (*models.SocialMediaModel, error) {
	socialMedia := models.SocialMediaModel{
		BaseModel: models.BaseModel{
			ID: socialMediaID,
		},
		UserID:         userID,
		Name:           dto.Name,
		SocialMediaUrl: dto.SocialMediaUrl,
	}

	result, err := s.socialMediaRepository.UpdateSocialMedia(&socialMedia)

	if err != nil {
		return &socialMedia, err
	}

	return result, nil
}

func (s *socialMediaService) DeleteSocialMedia(userID uint, socialMediaID uint) (*models.SocialMediaModel, error) {
	socialMedia := models.SocialMediaModel{
		BaseModel: models.BaseModel{
			ID: socialMediaID,
		},
	}

	result, err := s.socialMediaRepository.DeleteSocialMedia(&socialMedia)

	if err != nil {
		return &socialMedia, err
	}

	return result, nil
}

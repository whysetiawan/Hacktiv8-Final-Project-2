package services

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"
	"fmt"
)

type SocialMediaService interface {
	CreateSocialMedia(dto *dto.UpsertSocialMediaDto, userId uint) (*models.SocialMediaModel, error)
}

type socialMediaService struct {
	socialMediaRepository repositories.SocialMediaRepository
}

func NewSocialMediaService(r repositories.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{r}
}

func (s *socialMediaService) CreateSocialMedia(dto *dto.UpsertSocialMediaDto, userId uint) (*models.SocialMediaModel, error) {
	socialMedia := models.SocialMediaModel{
		UserId:         userId,
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

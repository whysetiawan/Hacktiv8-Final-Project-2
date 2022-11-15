package services

import (
	"errors"
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"
	"time"
)

type PhotoService interface {
	Create(input dto.InputPhoto, userID int64) (models.PhotoModelResponse, error)
	GetPhotoByID(ID int64) ([]models.GetPhotoResponse, error)
	UpdatePhoto(input dto.InputPhoto, photoId int64, userID int64) (models.PhotoModel, error)
	DeletePhoto(id, userid int64) (map[string]string, error)
}

type photoService struct {
	photoRepository repositories.PhotoRepository
}

func NewPhotoService(r repositories.PhotoRepository) *photoService {
	return &photoService{r}
}

func (ps photoService) Create(input dto.InputPhoto, userID int64) (models.PhotoModelResponse, error) {

	photo := models.PhotoModel{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()
	photo.UserId = uint(userID)

	newPhoto, err := ps.photoRepository.Save(photo)
	if err != nil {
		return newPhoto, err
	}

	result := models.PhotoModelResponse{
		ID:        newPhoto.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    userID,
		CreatedAt: photo.CreatedAt,
	}

	return result, nil
}

func (ps photoService) GetPhotoByID(ID int64) ([]models.GetPhotoResponse, error) {
	// var res []models.GetPhotoResponse

	photo, err := ps.photoRepository.GetPhotoById(ID)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

// func (ps photoService) UpdatePhoto()

func (ps photoService) UpdatePhoto(input dto.InputPhoto, photoId int64, userID int64) (models.PhotoModel, error) {

	photo, err := ps.photoRepository.FindByID(int64(photoId))
	if err != nil {
		return photo, err
	}

	if int64(photo.UserId) != userID {
		return photo, errors.New("you dont have access")
	}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl

	updatePhoto, err := ps.photoRepository.Update(photo)
	if err != nil {
		return updatePhoto, err
	}

	return updatePhoto, nil

}

func (ps photoService) DeletePhoto(id, userId int64) (map[string]string, error) {
	photo, err := ps.photoRepository.FindByID(id)
	if err != nil {
		return map[string]string{
			"en": "failed to check photo id",
		}, err
	}

	if photo.ID != uint(id) {
		return map[string]string{
			"en": "photo not found",
		}, err
	}

	if int64(photo.UserId) != userId {
		return map[string]string{
			"en": "you dont have access to delete photo",
		}, err
	}

	_, err = ps.photoRepository.Delete(photo, id)
	if err != nil {
		return map[string]string{
			"en":  "failed to delete data",
			"err": err.Error(),
		}, err
	}

	return map[string]string{
		"en": "Your photo has been successfully deleted",
	}, nil
}

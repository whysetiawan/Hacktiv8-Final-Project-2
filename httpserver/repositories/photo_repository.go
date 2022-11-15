package repositories

import (
	"final-project-2/httpserver/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Save(data models.PhotoModel) (models.PhotoModelResponse, error)
	GetPhotoById(ID int64) ([]models.GetPhotoResponse, error)
	FindByID(ID int64) (models.PhotoModel, error)
	Update(photo models.PhotoModel) (models.PhotoModel, error)
	Delete(photo models.PhotoModel, id int64) (int64, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (pr photoRepository) Save(data models.PhotoModel) (models.PhotoModelResponse, error) {
	var res models.PhotoModelResponse

	// q := `INSERT INTO Photos
	// (created_at, updated_at, deleted_at, title, photo_url, user_id)
	// VALUES(?, ?, ?, ?, ?, ?);
	// `

	// err := pr.db.Raw(q, data.CreatedAt, data.UpdatedAt, time.Now(), data.Title, data.PhotoUrl, data.UserId).Error
	err := pr.db.Create(&data).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pr photoRepository) GetPhotoById(ID int64) ([]models.GetPhotoResponse, error) {
	var res []models.GetPhotoResponse

	// err := pr.db.Where("user_id = ?", ID).Find(&res).Error
	// if err != nil {
	// 	return res, err
	// }

	// return res, nil

	err := pr.db.Preload("User").Raw(`select p.id, p.title, p.caption, p.photo_url , p.user_id, p.created_at , p.updated_at, u.username as username, u.email as email  from "Photo" p 
	inner join "User" u 
	on u.id = p.user_id 
	where p.user_id = ?`, ID).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pr photoRepository) FindByID(ID int64) (models.PhotoModel, error) {
	var user models.PhotoModel

	// err := pr.db.Where("id = ?", uint(ID)).Find(&user).Error
	err := pr.db.Raw(`select * from "Photo" where id =?`, ID).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (pr photoRepository) Update(photo models.PhotoModel) (models.PhotoModel, error) {
	err := pr.db.Updates(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (pr photoRepository) Delete(photo models.PhotoModel, id int64) (int64, error) {
	err := pr.db.Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

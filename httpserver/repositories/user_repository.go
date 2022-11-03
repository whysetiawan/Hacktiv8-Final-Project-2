package repositories

import (
	"final-project-2/httpserver/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.UserModel) (models.UserModel, error)
	DeleteTodo(todo models.UserModel) (models.UserModel, error)
	UpdateTodo(data models.UserModel, id int64) (models.UserModel, error)
}

type UserRepository interface {
	Register(user *models.UserModel) (*models.UserModel, error)
	Login(user *models.UserModel) (*models.UserModel, error)
	GetUsers() (*[]models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Login(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Find(&user).Where("email = ?", user.Email).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUsers() (*[]models.UserModel, error) {
	var users []models.UserModel
	err := r.db.Find(&users).Limit(10).Error

	if err != nil {
		return &users, err
	}

	return &users, nil
}

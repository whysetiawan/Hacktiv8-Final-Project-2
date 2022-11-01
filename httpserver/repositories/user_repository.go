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

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) CreateTodo(todo models.UserModel) (models.UserModel, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) DeleteTodo(todo models.UserModel) (models.UserModel, error) {
	err := r.db.Delete(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) UpdateTodo(data models.UserModel, id int64) (models.UserModel, error) {
	err := r.db.Model(&data).Where("todo_id", id).Updates(models.UserModel{
		// Status: true,
	})
	if err != nil {
		return data, err.Error
	}

	return data, nil
}

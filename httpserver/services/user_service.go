package services

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"
)

type TodoService interface {
	CreateTodo(dto dto.CreateTodoDto) (models.UserModel, error)
	DeleteTodo(id uint) error
	UpdateTodo(inputID int64, inputData dto.UpdateTodoDto) (map[string]string, error)
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(tr repositories.TodoRepository) *todoService {
	return &todoService{
		tr,
	}
}

func (s *todoService) CreateTodo(dto dto.CreateTodoDto) (models.UserModel, error) {
	todo := models.UserModel{
		// Name: dto.Name,
	}

	todo, err := s.todoRepository.CreateTodo(todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *todoService) DeleteTodo(id uint) error {
	todo := models.UserModel{
		// TodoId: id,
	}
	todo, err := s.todoRepository.DeleteTodo(todo)

	return err
}

func (s *todoService) UpdateTodo(inputID int64, inputData dto.UpdateTodoDto) (map[string]string, error) {
	_, err := s.todoRepository.UpdateTodo(models.UserModel{
		// Status: inputData.Status,
	}, inputID)

	if err != nil {
		return map[string]string{
			"en":    "fail to update status, please retry later",
			"id":    "update gagal",
			"error": err.Error(),
		}, err
	}

	return map[string]string{
		"en": "success",
		"id": "success update data",
	}, nil
}

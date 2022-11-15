package services

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"
	"fmt"
)

type CommentService interface {
	CreateComment(dto *dto.Comment, userId uint) (*models.CommentModel, error)
	GetComment() (*[]models.CommentModel, error)
	GetCommentByID(id uint) (models.CommentModel, error)
	UpdateComment(dto *dto.Comment, commentID uint, userID uint) (*models.CommentModel, error)
	DeleteComment(userID uint, commentID uint) (*models.CommentModel, error)
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(r repositories.CommentRepository) *commentService {
	return &commentService{r}
}

func (s *commentService) CreateComment(dto *dto.Comment, userId uint) (*models.CommentModel, error) {
	res := models.CommentModel{
		UserID:  userId,
		Message: dto.Message,
		PhotoID: dto.PhotoId,
	}
	fmt.Println(res)

	comment, err := s.commentRepository.CreateComment(&res)
	if err != nil {
		return &res, err
	}

	return comment, nil
}

func (s *commentService) GetComment() (*[]models.CommentModel, error) {
	return s.commentRepository.GetComment()
}

func (s *commentService) GetCommentByID(id uint) (models.CommentModel, error) {
	res, err := s.commentRepository.GetCommentByID(id)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s *commentService) UpdateComment(dto *dto.Comment, commentID uint, userID uint) (*models.CommentModel, error) {
	commentModel := models.CommentModel{
		BaseModel: models.BaseModel{
			ID: dto.ID,
		},
		Message: dto.Message,
	}

	commen, err := s.commentRepository.UpdateComment(&commentModel)
	if err != nil {
		return &commentModel, err
	}

	return commen, nil
}

func (s *commentService) DeleteComment(userID uint, commentID uint) (*models.CommentModel, error) {
	com := models.CommentModel{
		BaseModel: models.BaseModel{
			ID: commentID,
		},
	}

	res, err := s.commentRepository.DeleteComment(&com)
	if err != nil {
		return &com, err
	}

	return res, nil
}

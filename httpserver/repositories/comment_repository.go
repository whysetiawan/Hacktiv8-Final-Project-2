package repositories

import (
	"final-project-2/httpserver/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(commentCreate *models.CommentModel) (*models.CommentModel, error)
	GetComment() (*[]models.CommentModel, error)
	GetCommentByID(id uint) (models.CommentModel, error)
	UpdateComment(comment *models.CommentModel) (*models.CommentModel, error)
	DeleteComment(comment *models.CommentModel) (*models.CommentModel, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) CreateComment(commentCreate *models.CommentModel) (*models.CommentModel, error) {
	err := r.db.Create(&commentCreate).Error
	if err != nil {
		return commentCreate, err
	}
	return commentCreate, nil
}

func (r *commentRepository) GetComment() (*[]models.CommentModel, error) {
	var comment []models.CommentModel
	err := r.db.Find(&comment).Error

	if err != nil {
		return &comment, err
	}

	return &comment, nil
}

func (r *commentRepository) GetCommentByID(id uint) (models.CommentModel, error) {
	var idd models.CommentModel
	err := r.db.Find(id).Error
	if err != nil {
		return idd, err
	}

	return idd, nil
}

func (r *commentRepository) UpdateComment(comment *models.CommentModel) (*models.CommentModel, error) {
	err := r.db.Model(comment).Where(("id = ?"), comment.ID).Update("message", comment).Error
	if err != nil {
		return comment, err
	}
	return comment, err
}

func (r *commentRepository) DeleteComment(comment *models.CommentModel) (*models.CommentModel, error) {
	err := r.db.Model(comment).Where(("id = ?"), comment.ID).Delete(comment).Error
	if err != nil {
		return comment, err
	}
	return comment, err
}

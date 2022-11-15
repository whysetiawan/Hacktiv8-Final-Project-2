package controllers

import (
	"errors"
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetComment(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentController struct {
	commentService services.CommentService
}

func NewCommentController(
	commentService services.CommentService,
) *commentController {
	return &commentController{commentService}
}

// GetUsers godoc
// @Tags     Comment
// @Summary  Create Comment
// @Success  200 {object} utils.HttpSuccess[models.CommentModel]
// @Failure  401 {object} utils.HttpError
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /comment [post]
// @Security BearerAuth
func (c *commentController) CreateComment(ctx *gin.Context) {
	var dto dto.Comment
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	CommentCredential, isExist := ctx.Get("comment")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}
	CommentModel := CommentCredential.(models.CommentModel)

	res, err := c.commentService.CreateComment(&dto, CommentModel.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Comment Created", res))
}

// GetUsers godoc
// @Tags     Comment
// @Summary  get mutilple Comment
// @Success  200 {object} utils.HttpSuccess[[]models.CommentModel]
// @Failure  401 {object} utils.HttpError
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /comment [get]
// @Security BearerAuth
func (c *commentController) GetComment(ctx *gin.Context) {
	result, err := c.commentService.GetComment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Get All Success", result))
}

// GetUsers godoc
// @Tags     Comment
// @Summary  get Comment
// @Success  200 {object} utils.HttpSuccess[[]models.CommentModel]
// @Failure  401 {object} utils.HttpError
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /comment/:commentID [get]
// @Security BearerAuth
func (c *commentController) GetCommentByID(ctx *gin.Context) {
	idd := ctx.Param("commentID")
	tes, err := strconv.Atoi(idd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	result, err := c.commentService.GetCommentByID(uint(tes))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Get Comment by ID Success", result))
}

// UpdateUser godoc
// @Tags    Comment
// @Summary Update Comment
// @Param   user body     dto.Comment true "Update Comment Based On Token"
// @Success 200  {object} utils.HttpSuccess[dto.Comment]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router   /comment/:commentID [put]
// @Security BearerAuth
func (c *commentController) UpdateComment(ctx *gin.Context) {
	var dto dto.Comment
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Failed to bind photo request", err.Error()))
		return
	}

	idd := ctx.Param("commentID")
	commentID, err := strconv.Atoi(idd)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	commentCredental, isExist := ctx.Get("comment")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	commentModel := commentCredental.(models.CommentModel)
	res, err := c.commentService.UpdateComment(&dto, uint(commentID), commentModel.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Failed to Update Comment", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Your comment has been successfully updated", res))
}

// DeleteComment godoc
// @Tags    Comment
// @Summary delete current comment based on JWT
// @Success 200  {object} utils.HttpSuccess[dto.Comment]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router   /comment/:commentID [delete]
// @Security BearerAuth
func (c *commentController) DeleteComment(ctx *gin.Context) {
	idd := ctx.Param("commentID")
	commentID, err := strconv.Atoi(idd)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	commentCredential, isExist := ctx.Get("comment")
	commentModel := commentCredential.(models.CommentModel)

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	_, err = c.commentService.DeleteComment(commentModel.ID, uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	message := fmt.Sprintf("Comment %d has been deleted", commentModel.ID)
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess(message, struct{}{}))
}

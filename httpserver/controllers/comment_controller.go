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

// CreateComment godoc
// @Tags     Comment
// @Summary  Create Comment
// @Param   user body     dto.UpsertComment true "Create User DTO"
// @Success  200 {object} utils.HttpSuccess[models.CommentModel]
// @Failure  401 {object} utils.HttpError
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /comment [post]
// @Security BearerAuth
func (c *commentController) CreateComment(ctx *gin.Context) {

	var dto dto.CreateCommentDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}
	userModel := userCredential.(models.UserModel)

	res, err := c.commentService.CreateComment(&dto, userModel.ID)
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

// UpdateComment godoc
// @Tags    Comment
// @Summary Update Comment
// @Param   user body     dto.UpdateCommentDto true "Update Comment Based On Token"
// @Success 200  {object} utils.HttpSuccess[models.CommentModel]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router   /comment/:commentID [put]
// @Security BearerAuth
func (c *commentController) UpdateComment(ctx *gin.Context) {
	fmt.Println("UPDATING COMMENT")
	var dto dto.UpdateCommentDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	id := ctx.Param("commentID")
	commentID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)
	res, err := c.commentService.UpdateComment(&dto, uint(commentID), userModel.ID)
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

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	_, err = c.commentService.DeleteComment(userModel.ID, uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	message := fmt.Sprintf("Comment %d has been deleted", userModel.ID)
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess(message, struct{}{}))
}

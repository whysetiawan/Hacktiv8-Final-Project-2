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

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMediaByID(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(
	socialMediaService services.SocialMediaService,
) *socialMediaController {
	return &socialMediaController{socialMediaService}
}

// CreateSocialMedia godoc
// @Tags     Social Media
// @Summary  create a social media
// @Param    SocialMedia body     dto.UpsertSocialMediaDto true "Create Social Media DTO"
// @Success  201         {object} utils.HttpSuccess[models.SocialMediaModel]
// @Failure  400         {object} utils.HttpError
// @Failure  500         {object} utils.HttpError
// @Router   /socialmedias [post]
// @Security BearerAuth
func (c *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var dto dto.UpsertSocialMediaDto
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

	socialMedia, err := c.socialMediaService.CreateSocialMedia(&dto, userModel.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Social Media Created", socialMedia))
<<<<<<< HEAD
}

// GetSocialMedias godoc
// @Tags     Social Media
// @Summary  get all social medias based on user
// @Success  200 {object} utils.HttpSuccess[[]models.SocialMediaModel]
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /socialmedias [get]
// @Security BearerAuth
func (c *socialMediaController) GetSocialMedias(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	socialMedia, err := c.socialMediaService.GetSocialMedias(userModel.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Success Get All Social Medias", socialMedia))
}

// UpdateSocialMediaByID godoc
// @Tags     Social Media
// @Summary  create a user
// @Param    id          path     int                      true "Social Media ID"
// @Param    SocialMedia body     dto.UpsertSocialMediaDto true "Update Social Media Based On User"
// @Success  200         {object} utils.HttpSuccess[models.SocialMediaModel]
// @Failure  400         {object} utils.HttpError
// @Failure  500         {object} utils.HttpError
// @Router   /socialmedias/{id} [put]
// @Security BearerAuth
func (c *socialMediaController) UpdateSocialMediaByID(ctx *gin.Context) {
	var dto dto.UpsertSocialMediaDto
	err := ctx.BindJSON(&dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	idParam := ctx.Param("id")
	socialMediaID, err := strconv.Atoi(idParam)

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

	_, err = c.socialMediaService.UpdateSocialMedia(&dto, uint(socialMediaID), userModel.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Update User Success", dto))
}

// DeleteSocialMedia godoc
// @Tags     SocialMedia
// @Summary  Delete Social Media By ID
// @Param    id          path     int                      true "Social Media ID"
// @Success  200 {object} utils.HttpSuccess[string]
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /socialmedias/{id} [delete]
// @Security BearerAuth
func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {

	idParam := ctx.Param("id")
	socialMediaID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")
	userModel := userCredential.(models.UserModel)

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	_, err = c.socialMediaService.DeleteSocialMedia(userModel.ID, uint(socialMediaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	message := fmt.Sprintf("Social Media ID %d has been deleted", socialMediaID)
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess(message, struct{}{}))
=======
>>>>>>> 9608a17a1a02f56d3ed18d6aa91a31f0730a7a73
}

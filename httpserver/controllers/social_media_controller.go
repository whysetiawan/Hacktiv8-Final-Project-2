package controllers

import (
	"errors"
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
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
// @Tags    Social Media
// @Summary create a user
// @Param   user body     dto.UpsertSocialMediaDto true "Create User DTO"
// @Success 201  {object} utils.HttpSuccess[models.SocialMediaModel]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /socialmedias [post]
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

	_, err = c.socialMediaService.CreateSocialMedia(&dto, userModel.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("User Registered", &dto))
}

package controllers

import (
	"errors"
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetPhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
	authService  utils.AuthHelper
}

func NewPhotoController(
	photoService services.PhotoService,
	authService utils.AuthHelper,
) *photoController {
	return &photoController{photoService, authService}
}

func (pc photoController) CreatePhoto(ctx *gin.Context) {
	var param dto.InputPhoto

	err := ctx.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request json", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")
	log.Println(userCredential)
	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request credential", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)
	// userID := int64(userModel.ID)
	userID := int64(userModel.ID)
	newPhoto, err := pc.photoService.Create(param, userID)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("User Registered", newPhoto))
}

func (pc photoController) GetPhoto(ctx *gin.Context) {
	// var param dto.InputPhoto

	// err := ctx.BindJSON(&param)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request json", err.Error()))
	// 	return
	// }

	userCredential, isExist := ctx.Get("user")
	log.Println(userCredential)
	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request credential", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)
	// userID := int64(userModel.ID)
	userID := int64(userModel.ID)
	newPhoto, err := pc.photoService.GetPhotoByID(userID)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("User Registered", newPhoto))
}

func (pc photoController) UpdatePhoto(ctx *gin.Context) {
	var param dto.InputPhoto

	// ctx.Request.URL.Query().Get("photoId")

	photoId := ctx.Request.URL.Query().Get("photoId")
	id, _ := strconv.Atoi(photoId)
	err := ctx.BindJSON(&param)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request json", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")
	log.Println(userCredential)
	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request credential", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)
	// userID := int64(userModel.ID)
	userID := int64(userModel.ID)
	newPhoto, err := pc.photoService.UpdatePhoto(param, int64(id), userID)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("User Registered", newPhoto))
}
func (pc photoController) DeletePhoto(ctx *gin.Context) {

	photoId := ctx.Request.URL.Query().Get("photoId")
	id, _ := strconv.Atoi(photoId)

	userCredential, isExist := ctx.Get("user")
	log.Println(userCredential)
	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request credential", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)
	// userID := int64(userModel.ID)
	userID := int64(userModel.ID)
	newPhoto, err := pc.photoService.DeletePhoto(int64(id), userID)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("User Registered", newPhoto))
}

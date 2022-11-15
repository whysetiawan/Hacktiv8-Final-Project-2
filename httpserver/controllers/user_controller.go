package controllers

import (
	"errors"
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	authService utils.AuthHelper
}

func NewUserController(
	userService services.UserService,
	authService utils.AuthHelper,
) *userController {
	return &userController{userService, authService}
}

// Register godoc
// @Tags    User
// @Summary create a user
// @Param   user body     dto.UpsertUserDto true "Create User DTO"
// @Success 201  {object} utils.HttpSuccess[dto.UpsertUserDto]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /user/register [post]
func (c *userController) Register(ctx *gin.Context) {
	var dto dto.UpsertUserDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	_, err = c.userService.Register(&dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("User Registered", &dto))
}

// Login godoc
// @Tags    User
// @Summary login a user
// @Param   user body     dto.LoginDto true "Login User DTO"
// @Success 200  {object} utils.HttpSuccess[models.LoginResponse]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /user/login [post]
func (c *userController) Login(ctx *gin.Context) {
	var dto dto.LoginDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	user, err := c.userService.Login(&dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewHttpError("Invalid Credentials", err.Error()))
		return
	}

	accessToken, refreshToken, err := c.authService.GenerateToken(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Login Success", models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}))
}

// GetUsers godoc
// @Tags     User
// @Summary  get mutilple users
// @Success  200 {object} utils.HttpSuccess[[]models.UserModel]
// @Failure  401 {object} utils.HttpError
// @Failure  400 {object} utils.HttpError
// @Failure  500 {object} utils.HttpError
// @Router   /user [get]
// @Security BearerAuth
func (c *userController) GetUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Get All Success", users))
}

// UpdateUser godoc
// @Tags    User
// @Summary create a user
// @Param   user body     dto.UpsertUserDto true "Update User Based On Token"
// @Success 200  {object} utils.HttpSuccess[dto.UpsertUserDto]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router   /user [put]
// @Security BearerAuth
func (c *userController) UpdateUser(ctx *gin.Context) {
	var dto dto.UpsertUserDto
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
	_, err = c.userService.UpdateUser(&dto, &userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Update User Success", dto))
}

// DeleteUser godoc
// @Tags    User
// @Summary delete current user based on JWT
// @Success 200  {object} utils.HttpSuccess[dto.UpsertUserDto]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router   /user [delete]
// @Security BearerAuth
func (c *userController) DeleteUser(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")
	userModel := userCredential.(models.UserModel)

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	_, err := c.userService.DeleteUser(&userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	message := fmt.Sprintf("User ID %d has been deleted", userModel.ID)
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess(message, struct{}{}))
}

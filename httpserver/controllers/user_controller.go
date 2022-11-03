package controllers

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	authService services.AuthService
}

func NewUserController(
	userService services.UserService,
	authService services.AuthService,
) *userController {
	return &userController{userService, authService}
}

// Register godoc
// @Tags    User
// @Summary create a user
// @Param   user body     dto.RegisterDto true "Create User DTO"
// @Success 201  {object} utils.HttpSuccess[dto.RegisterDto]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /user/register [post]
func (c *userController) Register(ctx *gin.Context) {
	var dto dto.RegisterDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	dto.Password = string(hashedPassword)

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
	print(dto.Password)
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

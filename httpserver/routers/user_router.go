package routers

import (
	controllers "final-project-2/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController) *gin.RouterGroup {
	userRouter := route.Group("/user")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
	}
	return userRouter
}

package routers

import (
	controllers "final-project-2/httpserver/controllers"
	"final-project-2/httpserver/services"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController, authService services.AuthService) *gin.RouterGroup {
	userRouter := route.Group("/user")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
		// userRouter.GET("", middleware.JwtGuard(authService), userController.GetUsers)
	}
	return userRouter
}

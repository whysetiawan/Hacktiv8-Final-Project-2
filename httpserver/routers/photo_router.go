package routers

import (
	controllers "final-project-2/httpserver/controllers"
	"final-project-2/httpserver/middleware"
	"final-project-2/utils"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(route *gin.RouterGroup, photoController controllers.PhotoController, authService utils.AuthHelper) *gin.RouterGroup {
	photoRouter := route.Group("/photos")
	{
		// photoRouter.POST("register",  photoController.CreatePhoto)
		// photoRouter.POST("login", photoController.Login)
		// photoRouter.GET("", middleware.JwtGuard(authService), photoController.GetUsers)
		photoRouter.POST("", middleware.JwtGuard(authService), photoController.CreatePhoto)
		photoRouter.GET("", middleware.JwtGuard(authService), photoController.GetPhoto)
		photoRouter.PUT("", middleware.JwtGuard(authService), photoController.UpdatePhoto)
		photoRouter.DELETE("", middleware.JwtGuard(authService), photoController.DeletePhoto)
	}
	return photoRouter
}

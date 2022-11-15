package routers

import (
	controllers "final-project-2/httpserver/controllers"
	"final-project-2/httpserver/middleware"
	"final-project-2/utils"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(route *gin.RouterGroup, socialMediaController controllers.SocialMediaController, authService utils.AuthHelper) *gin.RouterGroup {
	socialMediaRouter := route.Group("/socialmedias")
	{
		socialMediaRouter.POST("", middleware.JwtGuard(authService), socialMediaController.CreateSocialMedia)
	}
	return socialMediaRouter
}

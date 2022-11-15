package routers

import (
	controllers "final-project-2/httpserver/controllers"
	"final-project-2/httpserver/middleware"
	"final-project-2/utils"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.RouterGroup, commentController controllers.CommentController, authService utils.AuthHelper) *gin.RouterGroup {
	commentRoutes := router.Group("/comment")
	{
		commentRoutes.POST("", middleware.JwtGuard(authService), commentController.CreateComment) //disini
		commentRoutes.GET("", middleware.JwtGuard(authService), commentController.GetComment)
		commentRoutes.GET("/:commentID", middleware.JwtGuard(authService), commentController.GetCommentByID) //disini
		commentRoutes.PUT("/:commentID", middleware.JwtGuard(authService), commentController.UpdateComment)
		commentRoutes.DELETE("/:commentID", middleware.JwtGuard(authService), commentController.DeleteComment)
	}
	return commentRoutes
}

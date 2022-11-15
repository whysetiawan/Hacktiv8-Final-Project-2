package main

import (
	"final-project-2/config"
	"final-project-2/docs"
	"final-project-2/httpserver/controllers"
	"final-project-2/httpserver/repositories"
	"final-project-2/httpserver/routers"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"log"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:3030
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Environment Variables not found")
	// }
	app := gin.Default()
	appRoute := app.Group("/api")
	db, err := config.Connect()
	if err != nil {
		log.Println("err db")
	}

	authService := utils.NewAuthHelper()

	userRepository := repositories.NewUserRepository(db)
	photoRepo := repositories.NewPhotoRepository(db)
	userService := services.NewUserService(userRepository)
	photoService := services.NewPhotoService(photoRepo)
	// authService := utils.NewAuthHelper()
	userController := controllers.NewUserController(userService, authService)
	photoController := controllers.NewPhotoController(photoService, authService)
	// todoRepository := repositories.NewTodoRepository(db)
	// todoService := services.NewTodoService(todoRepository)
	// todoController := controllers.NewTodoController(todoService)

	socialMediaRepository := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	routers.PhotoRouter(appRoute, photoController, authService)
	routers.UserRouter(appRoute, userController, authService)
	routers.SocialMediaRouter(appRoute, socialMediaController, authService)

	docs.SwaggerInfo.Title = "Hacktiv8 final-project-2 API"
	docs.SwaggerInfo.Description = "This is just a simple TODO List"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// log.Println("Generating Swagger")
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(path) // for example /home/user
	// cmd, err := exec.Command("swag", "fmt").Output()
	// log.Println(cmd)
	// log.Println("Swagger Generated")

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(":3000")
}

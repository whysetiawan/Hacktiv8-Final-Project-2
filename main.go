package main

import (
	"final-project-2/config"
	"final-project-2/docs"
	"final-project-2/httpserver/controllers"
	"final-project-2/httpserver/repositories"
	"final-project-2/httpserver/routers"
	"final-project-2/httpserver/services"
	"final-project-2/utils"

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
	db, _ := config.Connect()

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	socialMediaRepository := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	routers.UserRouter(appRoute, userController, authService)
	routers.SocialMediaRouter(appRoute, socialMediaController, authService)

	docs.SwaggerInfo.Title = "Hacktiv8 final-project-2 API"
	docs.SwaggerInfo.Description = "This is just a simple TODO List"
	docs.SwaggerInfo.Host = "hacktiv8-final-project-2-production.up.railway.app"
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
	app.Run()
}

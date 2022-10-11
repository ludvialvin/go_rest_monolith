package main

import (
	"fmt"
	"go_rest_monolith/config"
	"go_rest_monolith/config/database"
	"go_rest_monolith/utils"
	"go_rest_monolith/utils/middleware/rbac"
	"go_rest_monolith/routes"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "go_rest_monolith/docs"
)

// @title 						Golang REST API
// @version 					1.0
// @description 				This is a sample swagger for Fiber Go
// @termsOfService 				https://mirav.in/terms
// @contact.name 				API Support
// @contact.email 				ludvi.alvin.office@gmail.com
// @host 						localhost:3000
// @BasePath 					/
// @securityDefinitions.apikey 	TokenAuth
// @in 							header
// @name 						Authorization
// @description 				TokenAuth protects our entity endpoints
func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(logger.New())
	//app.Use(cors.New())

	routes.AuthRoutes(app)
	routes.RBACRoutes(app)
	routes.UserRoutes(app)

	rbac.InitRBAC()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}

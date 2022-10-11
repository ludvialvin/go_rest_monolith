package routes

import (
	"go_rest_monolith/controllers"
	"go_rest_monolith/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	r := app.Group("/user").Use(middleware.Auth)

	r.Get("/", controllers.GetUsers)
	r.Get("/:id", controllers.GetUser)
	r.Post("/", controllers.CreateUser)
	r.Patch("/:id", controllers.UpdateUser)
	r.Delete("/:id", controllers.DeleteUser)
}
package routes

import (
	"go_rest_monolith/controllers"
	
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", controllers.Signup)
	r.Post("/login", controllers.Login)
}

package routes

import (
	"go_rest_monolith/controllers"
	"go_rest_monolith/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func RBACRoutes(app fiber.Router) {
	r := app.Group("/rbac").Use(middleware.Auth)

	r.Get("/roles", controllers.GetRoles)
	r.Post("/roles", controllers.CreateRole)
	r.Delete("/roles/:role_id", controllers.DeleteRole)
	r.Get("/permissions", controllers.GetPermissions)
	r.Get("/permissions/:role_id", controllers.GetPermissionsByRoleId)
	r.Post("/permissions/create", controllers.SetRolePermission)
	r.Post("/permissions/delete", controllers.RemoveRolePermission)
}
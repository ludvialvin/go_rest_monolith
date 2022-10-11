package migration

import (
	"go_rest_monolith/config/database"
	"go_rest_monolith/types"
	"go_rest_monolith/utils/middleware/rbac"
)

func InitDefaultData() {
	db := database.Gorm2

	if !db.Migrator().HasTable(&types.User{}) {
		db.AutoMigrate(&types.User{})
	}

	countUser := int64(0)
	db.Model(types.User{}).Count(&countUser)
	if countUser == 0 {
		db.Create(&types.User{Name: "admin", Email: "admin@local.com", User_group_id: 1, Password: "$2a$10$6x81Z5jxljjGoZiPXcdEiOPKnPQEYetPxZkGOaLd/GCKXuBvwH7Vu"})
		db.Create(&types.User{Name: "user", Email: "user@local.com", User_group_id: 2, Password: "$2a$10$6x81Z5jxljjGoZiPXcdEiOPKnPQEYetPxZkGOaLd/GCKXuBvwH7Vu"})
	}

	rbac.Auth.CreateRole("ADMIN")
	rbac.Auth.CreateRole("USER")

	//create permission rbac
	rbac.Auth.CreatePermission("GetRoles")
	rbac.Auth.CreatePermission("CreateRole")
	rbac.Auth.CreatePermission("DeleteRole")
	rbac.Auth.CreatePermission("GetPermissions")
	rbac.Auth.CreatePermission("GetPermissionsByRoleId")
	rbac.Auth.CreatePermission("SetRolePermission")
	rbac.Auth.CreatePermission("RemoveRolePermission")

	//create permission users
	rbac.Auth.CreatePermission("GetUsers")
	rbac.Auth.CreatePermission("GetUser")
	rbac.Auth.CreatePermission("CreateUser")
	rbac.Auth.CreatePermission("UpdateUser")
	rbac.Auth.CreatePermission("DeleteUser")

	//assign permission roles admin
	rbac.Auth.AssignPermissions("ADMIN", []string{
		"GetRoles",
		"CreateRole",
		"DeleteRole",
		"GetPermissions",
		"GetPermissionsByRoleId",
		"SetRolePermission",
		"RemoveRolePermission",
		"GetUsers",
		"GetUser",
		"CreateUser",
		"UpdateUser",
		"DeleteUser",
	})
}

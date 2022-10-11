package models

import (
	"go_rest_monolith/types"
	"go_rest_monolith/config/database"
	"go_rest_monolith/utils/middleware/pagination"
	"go_rest_monolith/utils/middleware/rbac"
	"gorm.io/gorm"
	"strings"
	"errors"
)

func FindRoles(params types.Filter) types.Result {
	var roles []types.Roles
	db := database.Gorm2

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Table:	 "authority_roles",
		Select:  "id, name",
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: []string{params.SortBy+" "+params.SortDir},
		ShowSQL: false,
	}, &roles)

	statusCode := 200
	if len(roles) < 1 {
		statusCode = 404
	}

	res := types.Result{
    	Status: "success",
        StatusCode: statusCode,
        Data: paginator,
    }

	return res
}

func CreateRole(name string) types.MsgResponse {
	if name == "" {
    	res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Role name empty",
	    }

	    return res
    }

    roleName := strings.ToUpper(name)

    err := rbac.Auth.CreateRole(roleName)

    if err != nil {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Failed create role",
	    }

	    return res
	}

	res := types.MsgResponse{
    	Status: "success",
        StatusCode: 200,
        Message: name,
    }

    return res
}	

func GetRoleById(id int) *gorm.DB {
	return database.Gorm2.Model(&types.Authority_roles{}).Where("id = ?", id).First(&types.Authority_roles{})
}

func GetRoleByName(name string) *gorm.DB {
	return database.Gorm2.Model(&types.Authority_roles{}).Where("name = ?", name).First(&types.Authority_roles{})
}

func DeleteRole(id int) *gorm.DB {
	return database.Gorm2.Model(&types.Authority_roles{}).Where("id = ?", id).Delete(&types.Authority_roles{})
}

func FindPermissions(params types.Filter) types.Result {
	var permissions []types.Permissions
	db := database.Gorm2
	
	paginator := pagination.Paging(&pagination.Param{
		DB:      db.Table("authority_role_permissions a").Select("a.id, a.role_id, b.name as role_name, a.permission_id, c.name as permission_name").Joins("left join authority_roles b on b.id = a.role_id").Joins("left join authority_permissions c on c.id = a.permission_id").Find(&permissions),
		Table:	 "authority_role_permissions",
		Select:  "id, role_id, permission_id",
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: []string{params.SortBy+" "+params.SortDir},
		ShowSQL: false,
	}, &permissions)

	statusCode := 200
	if len(permissions) < 1 {
		statusCode = 404
	}

	res := types.Result{
    	Status: "success",
        StatusCode: statusCode,
        Data: paginator,
    }

	return res
}

func FindPermissionsByRoleId(params types.Filter, role_id int) types.Result {
	var permissions []types.Permissions
	db := database.Gorm2
	
	paginator := pagination.Paging(&pagination.Param{
		DB:      db.Table("authority_role_permissions a").Select("a.id, a.role_id, b.name as role_name, a.permission_id, c.name as permission_name").Joins("left join authority_roles b on b.id = a.role_id").Joins("left join authority_permissions c on c.id = a.permission_id").Where("a.role_id = ?",role_id).Find(&permissions),
		Table:	 "authority_role_permissions",
		Select:  "id, role_id, permission_id",
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: []string{params.SortBy+" "+params.SortDir},
		ShowSQL: false,
	}, &permissions)

	statusCode := 200
	if len(permissions) < 1 {
		statusCode = 404
	}

	res := types.Result{
    	Status: "success",
        StatusCode: statusCode,
        Data: paginator,
    }

	return res
}

func SetRolePermission(role_id int, permission_id int) types.MsgResponse {
	var roleName types.Authority_roles

	resRole := database.Gorm2.Model(types.Authority_roles{}).Where("id = ?", role_id).First(&roleName)

	if errors.Is(resRole.Error, gorm.ErrRecordNotFound) {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "User Group not found",
	    }

	    return res
	}

	var permissionName types.Authority_permissions

	resPermission := database.Gorm2.Model(types.Authority_permissions{}).Where("id = ?", permission_id).First(&permissionName)

	if errors.Is(resPermission.Error, gorm.ErrRecordNotFound) {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Permission not found",
	    }

	    return res
	}

	err := rbac.Auth.AssignPermissions(roleName.Name, []string{
		permissionName.Name,
	})

	if err != nil {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Failed set permission "+permissionName.Name+" for Role "+roleName.Name,
	    }

	    return res
	}

	res := types.MsgResponse{
    	Status: "success",
        StatusCode: 200,
        Message: "Success set permission "+permissionName.Name+" for Role "+roleName.Name,
    }

    return res
}

func RemoveRolePermission(role_id int, permission_id int) types.MsgResponse {
	var roleName types.Authority_roles

	resRole := database.Gorm2.Model(types.Authority_roles{}).Where("id = ?", role_id).First(&roleName)

	if errors.Is(resRole.Error, gorm.ErrRecordNotFound) {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "User Group not found",
	    }

	    return res
	}

	var permissionName types.Authority_permissions

	resPermission := database.Gorm2.Model(types.Authority_permissions{}).Where("id = ?", permission_id).First(&permissionName)

	if errors.Is(resPermission.Error, gorm.ErrRecordNotFound) {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Permission not found",
	    }

	    return res
	}

	err := rbac.Auth.RevokeRolePermission(roleName.Name, permissionName.Name)

	if err != nil {
		res := types.MsgResponse{
	    	Status: "failed",
	        StatusCode: 404,
	        Message: "Failed remove permission "+permissionName.Name+" for Role "+roleName.Name,
	    }

	    return res
	}

	res := types.MsgResponse{
    	Status: "success",
        StatusCode: 200,
        Message: "Success remove permission "+permissionName.Name+" for Role "+roleName.Name,
    }

    return res
}
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_rest_monolith/models"
	"go_rest_monolith/types"
	"go_rest_monolith/utils/middleware"
	"go_rest_monolith/utils/middleware/rbac"
	"gorm.io/gorm"
	"errors"
	"strconv"
)

func GetRoles(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "GetRoles")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	params := middleware.ParseQuery(c)
	res := models.FindRoles(params)

	return c.Status(res.StatusCode).JSON(res)
}	

func CreateRole(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "CreateRole")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	data := new(types.CreateRole)

    if err := c.BodyParser(data); err != nil {
        return err
    }

	res := models.CreateRole(data.Name)

	return c.Status(res.StatusCode).JSON(res)
}	

func DeleteRole(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "DeleteRole")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	id, _ := strconv.Atoi(c.Params("role_id"))

	if id == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Invalid ID")
	}

	err2 := models.GetRoleById(id).Error

	if errors.Is(err2, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 404,
			Message: "Role not found",
		})
	}

	err := models.DeleteRole(id).Error

	if err != nil {
		return c.Status(400).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 400,
			Message: "Delete role failed",
		})
	}

	return c.Status(200).JSON(&types.MsgResponse{
		Status: "success",
		StatusCode: 200,
		Message: "Role successfully deleted",
	})
}

func GetPermissions(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "GetPermissions")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	params := middleware.ParseQuery(c)

	res := models.FindPermissions(params)

	return c.Status(res.StatusCode).JSON(res)
}

func GetPermissionsByRoleId(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "GetPermissionsByRoleId")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	params := middleware.ParseQuery(c)

	role_id, _ := strconv.Atoi(c.Params("role_id"))

	res := models.FindPermissionsByRoleId(params,role_id)

	return c.Status(res.StatusCode).JSON(res)
}	

func SetRolePermission(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "SetRolePermission")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	data := new(types.CreateRolePermission)

    if err := c.BodyParser(data); err != nil {
        return err
    }

    if data.Role_id == "" {
    	return c.Status(404).JSON(fiber.Map{"status":"failed", "statusCode": 404, "error": "Role ID empty"})
    }

    if data.Permission_id == "" {
    	return c.Status(404).JSON(fiber.Map{"status":"failed", "statusCode": 404, "error": "Permission_id ID empty"})
    }

    role_id, _ := strconv.Atoi(data.Role_id)
    permission_id, _ := strconv.Atoi(data.Permission_id)

	res := models.SetRolePermission(role_id,permission_id)

	return c.Status(res.StatusCode).JSON(res)
}

func RemoveRolePermission(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "RemoveRolePermission")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	data := new(types.DeleteRolePermission)

    if err := c.BodyParser(data); err != nil {
        return err
    }

    if data.Role_id == "" {
    	return c.Status(404).JSON(fiber.Map{"status":"failed", "statusCode": 404, "error": "Role ID empty"})
    }

    if data.Permission_id == "" {
    	return c.Status(404).JSON(fiber.Map{"status":"failed", "statusCode": 404, "error": "Permission_id ID empty"})
    }

    role_id, _ := strconv.Atoi(data.Role_id)
    permission_id, _ := strconv.Atoi(data.Permission_id)

	res := models.RemoveRolePermission(role_id,permission_id)

	return c.Status(res.StatusCode).JSON(res)
}
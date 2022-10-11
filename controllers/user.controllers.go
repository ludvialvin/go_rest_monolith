package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_rest_monolith/models"
	"go_rest_monolith/types"
	"go_rest_monolith/utils"
	"go_rest_monolith/utils/middleware"
	"go_rest_monolith/utils/password"
	"go_rest_monolith/utils/middleware/rbac"
	"gorm.io/gorm"
	"errors"
	"strconv"
)

// @Tags 		User
// @Summary 	Get list
// @Description Get list users data
// @ID 			get-user
// @Produce  	json
// @Security 	TokenAuth
// @Param 		limit query int false "Limit"
// @Param 		page query int false "Page"
// @Param 		sortby query string false "Sort By"
// @Param 		sortdir query string false "Sort Dir asc/desc"
// @Success 	200 {object} types.UserResp
// @Failure 	401 {string} string "Not authorized"
// @Router 		/user [get]
func GetUsers(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "GetUsers")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	params := middleware.ParseQuery(c)

	res := models.FindUsers(params)

	return c.Status(res.StatusCode).JSON(res)
}

// @Tags 		User
// @Summary 	Get one
// @Description Get users data by user id
// @ID 			get-user-by-id
// @Produce  	json
// @Security 	TokenAuth
// @Param   	id      path   int     true  "User ID"
// @Success 	200 {object} types.UserResp
// @Failure 	400 {string} string "We need ID!!"
// @Failure 	401 {string} string "Not authorized"
// @Failure 	404 {string} string "Can not find ID"
// @Router 		/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "GetUser")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	id, _ := strconv.Atoi(c.Params("id"))

	if id == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Invalid ID")
	}

	res := models.FindUser(id)

	return c.Status(res.StatusCode).JSON(res)
}

// @Tags 		User
// @Summary 	Create
// @Description Create user
// @ID 			create-user
// @Produce  	json
// @Security 	TokenAuth
// @Param 		Body body string true "User Json" SchemaExample({\r\n"name":"User Name",\r\n"email":"user@example.com",\r\n"password":"123456"\r\n})
// @Success 	200 {object} types.MsgResponse
// @Failure 	401 {string} string "Not authorized"
// @Router 		/user [post]
func CreateUser(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "CreateUser")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	b := new(types.User)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := models.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &types.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	if err := models.CreateUser(user).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{"Status":"success","StatusCode":200})
}	

// @Tags 		User
// @Summary 	Update
// @Description Update user
// @ID 			update-user
// @Produce  	json
// @Security 	TokenAuth
// @Param   	id      path   int     true  "User ID"
// @Param 		email body string true "User Json" SchemaExample({\r\n"name":"User Name",\r\n"email":"user@example.com",\r\n"password":"123456"\r\n})
// @Success 	200 {object} types.MsgResponse
// @Failure 	400 {string} string "We need ID!!"
// @Failure 	401 {string} string "Not authorized"
// @Failure 	404 {string} string "Can not find ID"
// @Router 		/user/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "UpdateUser")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	id, _ := strconv.Atoi(c.Params("id"))
	b := new(types.User)

	if id == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Invalid ID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err2 := models.GetUserById(id).Error

	if errors.Is(err2, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 404,
			Message: "User not exist",
		})
	}

	user := &types.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	err := models.UpdateUser(id,user).Error

	if err != nil {
		return c.Status(400).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 400,
			Message: "Update user failed",
		})
	}

	return c.Status(200).JSON(&types.MsgResponse{
		Status: "success",
		StatusCode: 200,
		Message: "User successfully updated",
	})
}

// @Tags 			User
// @Summary 		Delete
// @Description 	Delete user
// @ID 				delete-user
// @Produce  		json
// @Security 		TokenAuth
// @Param   		id      path   int     true  "User ID"
// @Success 		200 {object} types.MsgResponse
// @Failure 		400 {string} string "We need ID!!"
// @Failure 		401 {string} string "Not authorized"
// @Failure 		404 {string} string "Can not find ID"
// @Router 			/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	chkPermission, _ := rbac.Auth.CheckRolePermission(c.Locals("USER_ROLE").(string), "DeleteUser")

	if chkPermission == false {
		return fiber.ErrUnauthorized
	}

	id, _ := strconv.Atoi(c.Params("id"))

	if id == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Invalid ID")
	}

	err2 := models.GetUserById(id).Error

	if errors.Is(err2, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 404,
			Message: "User not exist",
		})
	}

	err := models.DeleteUser(id).Error

	if err != nil {
		return c.Status(400).JSON(&types.MsgResponse{
			Status: "failed",
			StatusCode: 400,
			Message: "Delete user failed",
		})
	}

	return c.Status(200).JSON(&types.MsgResponse{
		Status: "success",
		StatusCode: 200,
		Message: "User successfully deleted",
	})
}
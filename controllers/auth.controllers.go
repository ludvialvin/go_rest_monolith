package controllers

import (
	"go_rest_monolith/config"
	"go_rest_monolith/models"
	"go_rest_monolith/types"
	"go_rest_monolith/utils"
	"go_rest_monolith/utils/jwt"
	"go_rest_monolith/utils/password"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"errors"
)

// @Tags        Auth
// @Summary 	Login
// @Description Login
// @ID 			login
// @Accept  	multipart/form-data
// @Produce  	json
// @Param 		email formData string true "email"
// @Param 		password formData string true "password"
// @Param 		X-API-KEY header string true "API Key"
// @Success 	200 {object} types.AuthResponse
// @Failure 	401 {string} string "Not authorized"
// @Router 		/auth/login [post]
func Login(ctx *fiber.Ctx) error {
	b := new(types.LoginDTO)

	apiKey := ctx.Get("X-API-KEY")

	if apiKey == "" {
		return fiber.ErrUnauthorized
	}

	if apiKey != config.X_API_KEY {
		return fiber.ErrUnauthorized
	}

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := models.FindUserByEmail(u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	return ctx.Status(200).JSON(&types.AuthResponse{
		Status: "success",
		StatusCode: 200,
		User: u,
		Token: t,
	})
}

// Signup service creates a user
func Signup(ctx *fiber.Ctx) error {
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	err := models.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &types.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	// Create a user, if error return
	if err := models.CreateUser(user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return ctx.Status(200).JSON(&types.AuthResponse{
		Status: "success",
		StatusCode: 200,
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Token: t,
	})
}
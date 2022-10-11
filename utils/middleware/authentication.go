package middleware

import (
	"go_rest_monolith/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go_rest_monolith/config/database"
	//"fmt"
)

// Auth is the authentication middleware
type Authority_role struct {
	Name string `json:"name"`
}

func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	// Spliting the header
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("USER_ID", user.ID)

	role := Authority_role{}
	database.Gorm2.Table("users").Select("authority_roles.name").Joins("left join authority_roles on authority_roles.id = users.user_group_id").Where("users.id = ?",user.ID).Limit(1).Find(&role)
	c.Locals("USER_ROLE", role.Name)

	return c.Next()
}

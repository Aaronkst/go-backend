package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Local Functions

// Exported Functions

func Auth(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if strings.Contains(auth, "Bearer") {
		token := strings.Split(auth, " ")[1]
		c.Locals("user", token)
	}
	return c.Next()
}

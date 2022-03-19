package main

import (
	auth "go-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", auth.Auth, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/", auth.Auth, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"data":   c.Locals("user"),
		})
	})

	app.Listen(":3000")
}

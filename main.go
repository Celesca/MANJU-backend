package main

import (
	"manju/backend/config/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	api := app.Group("/api/v1")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}

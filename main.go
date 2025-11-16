package main

import (
	"log"
	"manju/backend/config/database"

	"github.com/gofiber/fiber/v2"
	routes "manju/backend/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	api := app.Group("/api/v1")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.UserRoutes(api)

	log.Fatal(app.Listen(":8080"))
}

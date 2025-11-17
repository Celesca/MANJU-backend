package main

import (
	"log"
	"manju/backend/config/database"

	routes "manju/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	database.Connect()
	app := fiber.New()

	api := app.Group("/api")
	api.Get("/docs/*", swagger.HandlerDefault) // default swagger UI
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.UserRoutes(api)
	routes.VoiceRoutes(api)

	log.Fatal(app.Listen(":8080"))
}

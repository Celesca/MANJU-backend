package main

import (
	"log"
	"manju/backend/config/database"

	routes "manju/backend/routes"

	"github.com/gofiber/contrib/swagger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(swagger.New())
	api := app.Group("/api/v1")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.UserRoutes(api)

	log.Fatal(app.Listen(":8080"))
}

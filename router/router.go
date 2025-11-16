package router

import (
	"log"
	"manju/backend/handlers"
	"manju/backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRouter() *fiber.App {
	// Init DB
	dsn := "host=localhost user=postgres password=postgres dbname=manju_dev port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}

	// Auto-migrate
	if err := db.AutoMigrate(&repository.User{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	repo := repository.New(db)

	app := fiber.New()
	app.Use(logger.New())

	// Health
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber")
	})

	handlers.RegisterUserRoutes(app, repo)

	return app
}

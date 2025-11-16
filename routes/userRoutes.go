package routes

import (
	"manju/backend/repository"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, repo *repository.UserRepository) {

	router := app.Group("/v1/users/")
	router.Post("/", func(c *fiber.Ctx) error { return createUser(c, repo) })
	router.Get("/", func(c *fiber.Ctx) error { return listUsers(c, repo) })
	router.Get("/:id", func(c *fiber.Ctx) error { return getUser(c, repo) })
	router.Patch("/:id", func(c *fiber.Ctx) error { return updateUser(c, repo) })
	router.Delete("/:id", func(c *fiber.Ctx) error { return deleteUser(c, repo) })
}

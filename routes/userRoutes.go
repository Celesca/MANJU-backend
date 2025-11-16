package routes

import (
	"manju/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {

	router := app.Group("/v1/users/")
	router.Get("/me", UserController.GetUser())
}

var UserController = controllers.NewUserController()

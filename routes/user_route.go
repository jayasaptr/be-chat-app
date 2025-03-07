package routes

import (
	"chat-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
}

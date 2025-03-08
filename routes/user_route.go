package routes

import (
	"chat-app/controllers"
	"chat-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	api.Get("/profile", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{"message": "Profile accessed", "user_id": userID})
	})
}

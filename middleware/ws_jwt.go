package middleware

import (
	"chat-app/utils"

	"github.com/gofiber/fiber/v2"
)

func WebSocketAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Query("token")

		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid Token"})
		}

		// Simpan user ID ke dalam context
		c.Locals("user_id", claims["user_id"])

		// Lanjut ke WebSocket handler
		return c.Next()
	}
}

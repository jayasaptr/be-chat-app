package controllers

import (
	"chat-app/services"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := services.RegisterUser(body.Username, body.Email, body.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

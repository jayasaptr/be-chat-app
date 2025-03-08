package routes

import (
	"chat-app/handlers"
	"chat-app/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupWebSocketRoutes(app *fiber.App) {
	app.Use("/ws", middleware.WebSocketAuth()) // Tambahkan middleware JWT
	app.Get("/ws", websocket.New(handlers.ChatHandler))
}

package main

import (
	"chat-app/config"
	"chat-app/handlers"
	"chat-app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	handlers.InitWebSocket()

	routes.SetupWebSocketRoutes(app)
	routes.SetupUserRoutes(app)

	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))

}

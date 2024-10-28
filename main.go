package main

import (
	"github.com/etharrra/go-chat/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// Create views engine
	viewEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	// Static route and directory
	app.Static("/static/", "./static")

	// Create a "ping" handler to test the server
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// create new App Handler
	appHandler := handlers.NewAppHandler()

	// Add appHandler routes
	app.Get("/", appHandler.HandleGetIndex)

	// create new websocket
	server := NewWebSocketServer()
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		server.HandleWebSocket(c)
	}))

	go server.HandelMessage()

	// Start the http server
	app.Listen("127.0.0.1:8080")
}

package routes

import (
	"MessagingService/api/handlers"
	"MessagingService/pkg/message"
	"github.com/gofiber/fiber/v2"
)

// MessageRouter is the Router for GoFiber App
func MessageRouter(app fiber.Router, service message.Service) {
	app.Use("/websocket", handlers.GetMessage(service))
	app.Post("/message", handlers.SendMessage(service))
}

package main

import (
	"MessagingService/api/routes"
	"MessagingService/pkg/message"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	service := message.NewService()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.MessageRouter(app.Group("/"), service)

	app.Listen(":3000")

}

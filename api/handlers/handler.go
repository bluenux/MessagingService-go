package handlers

import (
	"MessagingService/pkg/message"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func GetMessage(service message.Service) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}

func SendMessage(service message.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	}
}

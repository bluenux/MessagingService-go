package handlers

import (
	"MessagingService/pkg/entities"
	"MessagingService/pkg/message"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

// TODO : add GetMessage
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

		payload := new(entities.Payload)
		if err := ctx.BodyParser(&payload); err != nil {
			return err
		}

		service.SendMessage(payload)
		return ctx.SendString("Hello World!")
	}
}

func RegistryDevice(service message.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		deviceToken := ctx.FormValue("token")
		log.Printf("device token : %v\n", deviceToken)
		if len(deviceToken) == 0 {
			//log.Println("error", err)
			_ = ctx.SendStatus(fiber.StatusBadRequest)
			return ctx.SendString("ERROR!")
		}

		service.RegistryDevice(deviceToken)
		return ctx.SendString("OK!!")
	}

}

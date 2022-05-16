package main

import (
	"MessagingService/api/routes"
	"MessagingService/api/utility"
	"MessagingService/pkg/message"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"log"
)

func prepareHTTPServer() *fiber.App {
	log.Println("application init")

	app := fiber.New()
	service := message.NewService()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.MessageRouter(app.Group("/"), service)

	return app
}

func main() {

	app := prepareHTTPServer()

	if utility.IsAWSLambda() {
		startOnAWS()
	} else {
		start(app)
	}
}

func start(app *fiber.App) {
	_ = app.Listen(":3000")
}

func lambdaHandler(_ context.Context, req events.APIGatewayProxyRequest, app *fiber.App) (events.APIGatewayProxyResponse, error) {

	adapter := fiberadapter.New(app)
	resp, err := adapter.ProxyWithContext(context.Background(), req)
	return resp, err
}

func startOnAWS() {
	lambda.Start(lambdaHandler)
}

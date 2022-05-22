package main

import (
	"MessagingService/api/env"
	"MessagingService/api/routes"
	"MessagingService/pkg/message"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"log"
)

var app *fiber.App

func prepareHTTPServer() {
	log.Println("application init")

	app = fiber.New()
	service := message.NewService()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.MessageRouter(app.Group("/"), service)
}

func main() {

	prepareHTTPServer()

	if env.IsAWSLambda() {
		startOnAWS()
	} else {
		start()
	}
}

func start() {
	_ = app.Listen(":3000")
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	adapter := fiberadapter.New(app)
	resp, err := adapter.ProxyWithContext(context.Background(), req)
	return resp, err
}

func startOnAWS() {
	lambda.Start(lambdaHandler)
}

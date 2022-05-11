package message

import (
	"MessagingService/pkg/entities"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"log"
)

type Service interface {
	GetMessage()
	SendMessage(payload *entities.Payload)
}

type service struct {
	client *messaging.Client
}

func (s service) GetMessage() {
	//TODO implement me
	panic("implement me")
}

func (s service) SendMessage(payload *entities.Payload) {
	SendToToken(s.client, payload.Title, payload.Body)
}

func NewService() Service {
	// Obtain a messaging.Client from the App.
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := firebaseApp.Messaging(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return &service{client: client}
}

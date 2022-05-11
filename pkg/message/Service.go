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
	RegistryDevice(token string)
}

type service struct {
	client     *messaging.Client
	deviceList []string
}

func (s service) GetMessage() {
	//TODO implement me
	panic("implement me")
}

func (s service) SendMessage(payload *entities.Payload) {
	log.Printf("device len : %v\n", len(s.deviceList))
	if len(s.deviceList) == 0 {
		log.Println("no device!!")
	}
	for _, element := range s.deviceList {
		SendToToken(s.client, element, payload.Title, payload.Body)
	}
}

func (s *service) RegistryDevice(token string) {
	s.deviceList = append(s.deviceList, token)
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

	return &service{client: client, deviceList: []string{}}
}

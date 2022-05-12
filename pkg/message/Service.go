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
	RegistryDevice(token string) bool
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
		SendToToken(s.client, element, payload)
	}
}

func (s *service) RegistryDevice(token string) bool {
	if !s.isNewToken(token) {
		return false
	}
	if !s.isValidToken(token) {
		return false
	}

	s.deviceList = append(s.deviceList, token)
	log.Printf("added token : %v\n", token)

	return true
}

func (s *service) isNewToken(token string) bool {
	for _, element := range s.deviceList {
		if token == element {
			log.Printf("already registed!! : %v\n", token)
			return false
		}
	}
	return true
}

func (s *service) isValidToken(token string) bool {
	_, err := ValidateToken(s.client, token)
	if err != nil {
		log.Printf("invalid token : %v, error : %v\n", token, err)
		return false
	}
	return true
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

package message

import (
	"MessagingService/api/store"
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
	tokenStore store.TokenStore
}

func (s service) GetMessage() {
	//TODO implement me
	panic("implement me")
}

func (s service) SendMessage(payload *entities.Payload) {
	allToken := s.tokenStore.All()
	tokenCount := len(allToken)

	log.Printf("device len : %v\n", tokenCount)
	if tokenCount == 0 {
		log.Println("no device!!")
	}
	for _, element := range allToken {
		SendToToken(s.client, element, payload)
	}
}

func (s *service) RegistryDevice(token string) bool {
	if !s.isValidToken(token) {
		log.Println("invalid token!")
		return false
	}

	success := s.tokenStore.Set(token)
	log.Printf("token result : %v\n", success)

	return success
}

func (s *service) isValidToken(token string) bool {
	log.Println("token checking...!")
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

	tokenStore := store.NewStore()

	return &service{client: client, tokenStore: tokenStore}
}

package message

import (
	"MessagingService/pkg/entities"
	"context"
	"firebase.google.com/go/v4/messaging"
	"log"
)

func SendToToken(client *messaging.Client, token string, payload *entities.Payload) {

	// This registration token comes from the client FCM SDKs.
	registrationToken := token

	// See documentation on defining a message payload.
	message := &messaging.Message{
		/*Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},*/
		Data: map[string]string{
			"title":     payload.Title,
			"body":      payload.Body,
			"channelID": payload.ChannelID,
		},
		Token: registrationToken,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}
	log.Printf("request payload : %v\n", message)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Println(err)
		return
	}
	// Response is a message ID string.
	log.Printf("Successfully sent message: %v\n", response)
}

func ValidateToken(client *messaging.Client, token string) (string, error) {

	// This registration token comes from the client FCM SDKs.
	registrationToken := token

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Token: registrationToken,
	}
	log.Printf("request payload : %v\n", message)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.SendDryRun(context.Background(), message)
	log.Printf("response : %v\n", response)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Response is a message ID string.
	log.Printf("Successfully sent message: %v\n", response)

	return response, nil

}

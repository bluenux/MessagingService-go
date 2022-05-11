package message

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"log"
)

func SendToToken(client *messaging.Client, token string, title string, body string) {

	// This registration token comes from the client FCM SDKs.
	registrationToken := token

	// See documentation on defining a message payload.
	message := &messaging.Message{
		/*Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},*/
		Data: map[string]string{
			"title": title,
			"body":  body,
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

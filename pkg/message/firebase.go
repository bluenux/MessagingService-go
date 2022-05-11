package message

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"log"
)

func SendToToken(client *messaging.Client, title string, body string) {

	// This registration token comes from the client FCM SDKs.
	// TODO : dynamic token
	registrationToken := "token"

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
	fmt.Println("request payload : ", message)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

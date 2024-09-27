package services

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type MessagingService struct {
	client *messaging.Client
	token  string
}

func NewMessagingService(app *firebase.App, token string) *MessagingService {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	return &MessagingService{client: client, token: token}
}

func (s *MessagingService) SendMessage(title string, body string) error {
	ctx := context.Background()
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: s.token,
	}

	_, err := s.client.Send(ctx, message)
	if err != nil {
		log.Printf("error sending message: %v\n", err)
		return err
	}

	fmt.Println("Message sent successfully:", title, " - ", body)
	return nil
}

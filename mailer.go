package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
)

type mailer struct {
	client *sendgrid.SGClient
}

type Mailer interface {
	SendNotification(channel *Channel, datapoint *Datapoint) error
}

func NewMailer(username string, password string) Mailer {
	sg := sendgrid.NewSendGridClient(username, password)
	return &mailer{sg}
}

func (mailer *mailer) SendNotification(channel *Channel, datapoint *Datapoint) error {
	if len(channel.Email) == 0 {
		log.Printf("No notification address for channel %v", channel.Name)
		return nil
	}
	log.Printf("Sending notification for %v to %v", channel.Name, channel.Email)

	// TODO: Stop hardcoding.
	subject := fmt.Sprintf("[beerserver] Over limit: %f/26.0 @%s", datapoint.Value, channel.Name)
	text := fmt.Sprintf("Hurry up and cool down your beer!")

	mail := sendgrid.NewMail()
	mail.AddTo(channel.Email)
	mail.SetSubject(subject)
	mail.SetText(text)
	mail.SetFrom("notification@beerserver.local")

	err := mailer.client.Send(mail)
	if err != nil {
		log.Printf("Failed to send mail to %v", channel.Email)
	}
	return err
}

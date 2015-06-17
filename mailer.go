package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
)

type mailer struct {
	client *sendgrid.SGClient
}

type Mailer interface {
	SendNotification(to string, channel *Channel, datapoint *Datapoint) error
}

func NewMailer(username string, password string) Mailer {
	sg := sendgrid.NewSendGridClient(username, password)
	return &mailer{sg}
}

func (mailer *mailer) SendNotification(to string, channel *Channel, datapoint *Datapoint) error {
	// TODO: Stop hardcoding.
	subject := fmt.Sprintf("[beerserver] Over limit: %f/26.0 @%s", datapoint.Value, channel.Name)
	text := fmt.Sprintf("Hurry up and cool down your beer!")

	mail := sendgrid.NewMail()
	mail.AddTo(to)
	mail.SetSubject(subject)
	mail.SetText(text)
	mail.SetFrom("notification@beerserver.local")

	return mailer.client.Send(mail)
}

package main

import (
	"time"

	"github.com/tsawler/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {

	go func() {
		for {
			msg := <-app.MailChan
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTP
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectionTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, "Hello, <strong>world</strong>")

	err = email.Send(client)
}

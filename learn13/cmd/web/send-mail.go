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
}

package main

import (
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
}

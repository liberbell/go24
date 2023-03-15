package main

func listenForMail() {

	go func() {
		for {
			msg := <-app.MailChan
		}
	}()
}

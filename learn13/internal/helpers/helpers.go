package helpers

import "github.com/tsawler/bookings/internal/config"

var app *config.AppConfig

func NewHandlers(a *config.AppConfig) {
	app = a
}

package helpers

import (
	"net/http"

	"github.com/tsawler/bookings/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {

}

func ServerError(w http.ResponseWriter, err error) {

}

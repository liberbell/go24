package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tsawler/go-course/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandleFunc())

	return mux
}

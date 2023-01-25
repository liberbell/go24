package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("")
}

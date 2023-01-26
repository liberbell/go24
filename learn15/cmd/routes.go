package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tsawler/websockets-course/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandleFunc(handlers.SendData))

	return mux
}

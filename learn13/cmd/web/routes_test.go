package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/bookings/internal/config"
)

func testRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux: %T", v))

	}
}

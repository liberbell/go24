package dbrepo

import "github.com/tsawler/bookings/internal/config"

type postgresDBRepo struct {
	App *config.AppConfig
}

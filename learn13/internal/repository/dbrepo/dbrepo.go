package dbrepo

import (
	"database/sql"

	"github.com/tsawler/bookings/internal/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

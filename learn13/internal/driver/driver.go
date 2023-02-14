package driver

import (
	"database/sql"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {

}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("dbx", dsn)
}

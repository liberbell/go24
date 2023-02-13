package main

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// DatabaseURL := "postgres://dbmaster:@localhost:5432/test_connect"
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=dbmaster password=")
	// _, err := pgx.Connect(context.Background(), os.Getenv(DatabaseURL))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database.")

	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database.")
		return
	}
	log.Println("Pinged database.")

	err = getAllRows()
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	return nil
}

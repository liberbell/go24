package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	DatabaseURL := "postgres://dbmaster:@localhost:5432/test_connect"
	// _, err := pgx.Connect("pgx", "host=localhost port=5432 dbname=test_coneect user=dbmaster password")
	_, err := pgx.Connect(context.Background(), os.Getenv(DatabaseURL))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to conncet: %v\n", err))
	}
}

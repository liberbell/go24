package main

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	_, err := pgx.Connect("pgx", "host=localhost port=5432 dbname=test_coneect user=dbmaster password")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to conncet: %v\n", err))
	}
}

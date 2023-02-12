package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_coneect user=dbmaster password")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to conncet: %v\n", err))
	}
}

package main

import "database/sql"

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_coneect user=dbmaster password")
}

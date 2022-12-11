package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
	})
}

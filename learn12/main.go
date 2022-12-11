package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func main() {
	// fmt.Println("Hello World")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

const portNumber := ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
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
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}

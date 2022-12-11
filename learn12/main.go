package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
	}
}

func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannnot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil

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
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

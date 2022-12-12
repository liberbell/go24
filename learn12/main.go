package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
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

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

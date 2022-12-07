package main

import (
	"log"

	"github.com/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}

func main() {
	// PrintText("Hi")
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

// func PrintText(s string) {
// 	log.Println(s)
// }

package main

import "github.com/helpers"

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}

func main() {
	// PrintText("Hi")
	intChan := make(chan int)
}

// func PrintText(s string) {
// 	log.Println(s)
// }

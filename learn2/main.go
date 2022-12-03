package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
}

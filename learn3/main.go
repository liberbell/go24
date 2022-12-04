package main

import (
	"log"
	"time"
)

var s = "seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthday time.Time

type User struct {
	FirstName string
	LastName  string
	PhoneNaem string
	Age       int
	Birthday  time.Time
}

func main() {
	var s2 = "six"
	// s := "eight"

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s say from the saySomething func is ", s)
	return s3, "world"
}

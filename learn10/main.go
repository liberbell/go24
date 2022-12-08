package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Eric,
			"last_name": "Crapton,
			"hair_color": "grya",
			"has_dog": true
		},
		{
			"first_name": "Alex",
			"last_name": "hepp",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("Unmarshalled: %v", unmarshalled)

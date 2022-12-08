package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Eric",
			"last_name": "Crapton",
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

	var mySlice []Person
	var m1 Person
	m1.FirstName = "George"
	m1.LastName = "Harrison"
	m1.HairColor = "brown"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Bob"
	m2.LastName = "Marry"
	m2.HairColor = "white"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}

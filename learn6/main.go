package main

import (
	"log"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"
	// var firstLine = "Once upon a midnight dreary"
	// firstLine = "x"
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Bob", "Mary", "bob.mary@example.com", 54})
	users = append(users, User{"Elton", "John", "elton.john@example.com", 78})
	users = append(users, User{"Eric", "Crapton", "eric.crapton@example.com", 72})
	users = append(users, User{"Alex", "Hep", "alex.hep@example.com", 36})

	for _, l := range users {
		// log.Println(animalType, animal)
		// log.Println(i, ":", l)
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)

	}
}

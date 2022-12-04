package main

import (
	"log"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	animals := make(map[string]string)
	animals["dog"] = "Fido"

	for _, animal := range animals {
		log.Println(animal)
	}
}

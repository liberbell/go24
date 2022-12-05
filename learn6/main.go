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
	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"

	for i, l := range firstLine {
		// log.Println(animalType, animal)
		log.Println(i, ":", l)

	}
}

package main

import (
	"log"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}

package main

import "log"

func main() {
	// var isTrue bool

	// isTrue = false

	// if isTrue == true {
	// 	log.Println("isTrue is ", isTrue)
	// } else {
	// 	log.Println("isTrue is ", isTrue)
	// }

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("myNum is greater than 99 and isTrue is set to True")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("1")
	// } else if myNum == 101 && isTrue {
	// 	log.Println("2")
	// } else if myNum > 1000 && isTrue == false {
	// 	log.Println("3")
	// }

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")
	case "dog":
		log.Println("Cat is set to dog")
	}
}

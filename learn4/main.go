package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstname() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Elton"

	myVar2 := myStruct{
		FirstName: "Alex",
	}

	log.Println("myvar is set to ", myVar.FirstName)
	log.Println("myvar2 is set to ", myVar2.FirstName)
}

package main

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Elton"

	myVar2 := myStruct{
		FirstName: "Alex",
	}
}

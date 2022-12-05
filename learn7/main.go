package main

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name         string
	Color        string
	NumberOfLegs int
}

func main() {
	dog := Dog{
		Name:  "George",
		Breed: "German Shephered",
	}
}

package main

import "fmt"

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
	PrintInfo(dog)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal syas ", a.Says(), "and has", a.NumberOfLegs(). "legs.")
}

package main

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Eric,
			"last_name": "Crapton,
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

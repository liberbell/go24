package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"
	log.Panicln(s)
}

func saySomething(s string) (string, string) {
	return s, "world"
}

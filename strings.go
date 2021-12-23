package main

import "fmt"

/*
you can assign block of var like this way
better then storing them one by one!!
*/

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	fmt.Printf("%v, %T", actorName, actorName)
}

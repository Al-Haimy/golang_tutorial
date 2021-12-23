package main

import (
	"fmt"
)

var i int = 72

/*
the same variable can't be decleared in same
scope twice.
and if there is any new variable should be used
referance varaible shadwoing
*/
func main() {
	fmt.Println(i)
	var i int = 43
	fmt.Println(i)
}

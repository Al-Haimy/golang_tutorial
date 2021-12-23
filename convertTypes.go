package main

/*
# Variable declaaration
	- var foo int
	- var foo int = 42
	- foo := 42
# Can't redeclare variables, but can shadow them
# All variables must be used
# Visibility
	- lower case first letter for package scope
	- upper case first letter to export
	- no private scope
# Naming conventions
	- Pascal or camelCase
		* Caitalize acronyms(HTTP, URL)
	- As short as reasonable
		* longer names for longer lives
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j string
	// j = string(j) will not work but it can wokr with int() or float32()
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
	s := 3.43232
	j = strconv.FormatFloat(s, 'G', -1, 64)
	fmt.Printf("%v, %T\n", j, j)
	x := 3
	dd := float32(x) + 0.5
	fmt.Printf("%v, %T\n", dd, dd)

}

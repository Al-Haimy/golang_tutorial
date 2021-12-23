package main

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

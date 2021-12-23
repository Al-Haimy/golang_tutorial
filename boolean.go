package main

import (
	"fmt"
)

func main() {
	// var n bool = false
	n := 1 == 1 // true
	m := 1 == 2 // false
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	// Numeric Types
	//int
	var n2 uint16 = 43
	fmt.Printf("%v, %T\n", n2, n2)
	//arthematic
	a := 10 // 1010
	b := 3  // 0011
	// var b int8 = 3 most operatoin wont work unless they are both the same type
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // deviding two int won't give you floating
	fmt.Println(a % b)
	// operators

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

}

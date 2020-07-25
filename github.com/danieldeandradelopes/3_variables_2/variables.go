package main

import "fmt"

var (
	// Address = public (all packages) != address=private
	Address string // default
	// Tellphone = public (all packages) != tellphone=private
	Tellphone string  // default ""
	age       int     // default 0
	eat       bool    //default false
	value     float64 // default 0.00)
	words     rune
)

func main() {

	// sugar syntax (to declare a variables)
	test := "My value"

	fmt.Printf("address: %s\r\n", Address)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("eat: %v\r\n", eat)
	fmt.Printf("words: %v\r\n", words)
	fmt.Printf("test: %v\r\n", test)
}

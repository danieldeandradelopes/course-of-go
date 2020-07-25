package main

import "fmt"

var address string // default ""
var tellphone = "9999-9999"
var age int       // default 0
var eat bool      //default false
var value float64 // default 0.00
var words rune

func main() {
	fmt.Printf("address: %s\r\n", address)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("eat: %v\r\n", eat)
}

package main

import (
	"fmt"

	"4_packages/packages/prefix"
)

// UserName is public
var UserName = "Daniel"

func main() {
	fmt.Printf("Username is: %s\r\n", UserName)
	fmt.Printf("Prefix of capital: %d\r\n", prefix.Capital)

}

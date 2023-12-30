package main

import (
	"fmt"
)

func main() {
	switch i := 1; i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4, 5, 6, 7, 8, 9: // multiple cases
		fmt.Println("Four through Nine")
	default:
		fmt.Println("Unknown Number")
	}

	var s interface{} = "Hello"

	switch s.(type) {
	case string:
		fmt.Println("s is a string")
	case int:
		fmt.Println("s is an int")
	default:
		fmt.Println("Unknown type")
	}
}

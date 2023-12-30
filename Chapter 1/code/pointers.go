package main

import "fmt"

func main() {
	data := 20
	ptr := &data

	fmt.Println("Value stored at variable data = ", data)
	fmt.Println("Address of variable data = ", &data)
	fmt.Println("Value stored at variable ptr = ", ptr)
	fmt.Println("Value stored at variable *ptr = ", *ptr)

}

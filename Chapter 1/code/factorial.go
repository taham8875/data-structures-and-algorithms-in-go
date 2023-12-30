package main

import "fmt"

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println("The factorial of 5 is", factorial(5)) // output: "The factorial of 5 is 120"
}

package main

import "fmt"

// Take in a number from 1 through 1000, print Fizz if the number is divisible by 3, print Buzz if the number is divisible by 5, or FizzBuzz if the number is divisible by both. If none of these conditions are true, print nothing.

func main() {
	var i int
	fmt.Scanf("%d\n", &i)

	switch {
	case i%3 == 0 && i%5 == 0:
		fmt.Println("FizzBuzz")
	case i%3 == 0:
		fmt.Println("Fizz")
	case i%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println("")
	}

}

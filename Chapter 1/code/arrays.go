package main

import "fmt"

func main() {
	var array [10]int
	fmt.Println("The value of array is", array) // output: The value of array is [0 0 0 0 0 0 0 0 0 0]
	for i, _ := range array {
		array[i] = i
	}
	fmt.Println("The value of array is", array)       // output: The value of array is [0 1 2 3 4 5 6 7 8 9]
	fmt.Println("The length of array is", len(array)) // output: The length of array is 10
}

// Write a method, which will search a list for some given value.

package main

import "fmt"

func sequentialSearch(list []int, key int) int {
	for index, value := range list {
		if value == key {
			return index
		}
	}
	return -1
}

func main() {
	list := []int{-5, 30, 17, -26, 2, 42, 0, 12, 9, 5}
	key := 12
	fmt.Println("The value of key is", key)                         // output: The value of key is 3
	fmt.Println("The index of key is", sequentialSearch(list, key)) // output: The index of key is 7
}

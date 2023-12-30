// Write a method that will return the sum of all the elements of the integer list, given list as an input argument.

package main

import "fmt"

func sumList(list []int) int {
	sum := 0

	for _, value := range list {
		sum += value
	}

	return sum
}

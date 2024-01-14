package main

import "fmt"

func main() {
	var arr = []int{8, 2, 10, 9, 7, 5, 3, 1, 6, 4, 0}
	fmt.Println("Initial array is:", arr)
	InsertionSort(arr)
	fmt.Println("Sorted array is: ", arr)
}

func InsertionSort(array []int) {
	var temp, i, j int
	for i = 1; i < len(array); i++ {
		temp = array[i]
		for j = i; j > 0; j-- {
			if temp < array[j-1] {
				array[j] = array[j-1]
			} else {
				break
			}
		}
		array[j] = temp
	}
}

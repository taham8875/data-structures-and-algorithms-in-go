package main

import "fmt"

func main() {
	var arr = []int{8, 2, 10, 9, 7, 5, 3, 1, 6, 4, 0}
	fmt.Println("Initial array is:", arr)
	SelectionSort(arr)
	fmt.Println("Sorted array is: ", arr)
}

func SelectionSort(array []int) {
	var maximumElementIndex int
	for i := 0; i < len(array)-1; i++ {
		maximumElementIndex = 0
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[maximumElementIndex] {
				maximumElementIndex = j
			}
		}
		array[len(array)-i-1], array[maximumElementIndex] = array[maximumElementIndex], array[len(array)-i-1]
	}
}

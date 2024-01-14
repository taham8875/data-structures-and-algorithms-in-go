package main

import "fmt"

func main() {
	var arr = []int{8, 2, 10, 9, 7, 5, 3, 1, 6, 4, 0}
	fmt.Println("Initial array is:", arr)
	BubbleSort(arr)
	fmt.Println("Sorted array is: ", arr)
}

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

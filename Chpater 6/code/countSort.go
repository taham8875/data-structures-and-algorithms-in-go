package main

import "fmt"

func main() {
	var arr = []int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}
	fmt.Println("Initial array is:", arr)
	CountSort(arr, 21, 28)
	fmt.Println("Sorted array is: ", arr)
}

func CountSort(array []int, min int, max int) {
	var countArray = make([]int, max-min+1) // Go automatically initializes the array with 0s
	for i := 0; i < len(array); i++ {
		countArray[array[i]-min]++
	}
	var index int
	for i := 0; i < len(countArray); i++ {
		for countArray[i] > 0 {
			array[index] = i + min
			index++
			countArray[i]--
		}
	}
}

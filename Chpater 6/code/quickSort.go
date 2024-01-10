package main

import "fmt"

func main() {
	var arr = []int{8, 2, 10, 9, 7, 5, 3, 1, 6, 4, 0}
	fmt.Println("Initial array is:", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array is: ", arr)
}

func QuickSort(array []int, low int, high int) {
	if low > high || high < 0 || low < 0 || len(array) == 0 {
		return
	}
	pivot := Partition(array, low, high)
	QuickSort(array, low, pivot-1)
	QuickSort(array, pivot+1, high)
}

func Partition(array []int, low int, high int) int {
	pivot := array[high] // pivot is the last element, can be assumed to any other element
	i := low
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return i
}

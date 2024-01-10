package main

import "fmt"

func main() {
	var arr = []int{8, 2, 10, 9, 7, 5, 3, 1, 6, 4, 0}
	fmt.Println("Initial array is:", arr)
	MergeSort(arr)
	fmt.Println("Sorted array is: ", arr)

}

func MergeSort(array []int) {
	if len(array) <= 1 { // base case
		return
	}

	middle := len(array) / 2

	leftArray := make([]int, middle)
	rightArray := make([]int, len(array)-middle)

	copy(leftArray, array[:middle])
	copy(rightArray, array[middle:])

	MergeSort(leftArray)
	MergeSort(rightArray)

	Merge(leftArray, rightArray, array)
}

func Merge(leftArray []int, rightArray []int, array []int) {
	var i, j, k int

	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < len(leftArray) { // if rightArray is empty and there are still elements in leftArray
		array[k] = leftArray[i]
		i++
		k++
	}

	for j < len(rightArray) { // if leftArray is empty and there are still elements in rightArray
		array[k] = rightArray[j]
		j++
		k++
	}
}

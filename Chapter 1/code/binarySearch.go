package main

import "fmt"

func binarySearch(list []int, key int) int {
	low := 0
	high := len(list) - 1

	mid := (low + high) / 2

	for low <= high {
		mid = (low + high) / 2
		if key == list[mid] {
			return mid
		} else if key > list[mid] {
			low = mid + 1
		} else if key < list[mid] {
			high = mid
		}
	}

	return -1

}

func main() {
	// binarySearch works only on sorted lists
	list := []int{-26, -5, 0, 2, 5, 9, 12, 17, 30, 42}
	key := 12
	fmt.Println("The value of key is", key)                     // output: The value of key is 12
	fmt.Println("The index of key is", binarySearch(list, key)) // output: The index of key is 6
}

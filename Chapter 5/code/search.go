package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func GenerateRandomArrayOfIntegers(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(200)-100)
	}
	return array
}

func LinearSearchUnsortedData(array []int, key int) int {
	for i, v := range array {
		if v == key {
			return i
		}
	}
	return -1
}

func LinearSearchSortedData(array []int, key int) int {
	for i, v := range array {
		if v == key {
			return i
		} else if v > key { // If we've gone past where the key is, stop searching, no need to continue
			return -1
		}
	}
	return -1
}

func BinarySearch(array []int, key int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] < key {
			low = mid + 1
		} else if array[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	numbers := GenerateRandomArrayOfIntegers(10)
	key := numbers[rand.Intn(len(numbers))]
	fmt.Println("Unsorted: ", numbers)
	fmt.Println("Key: ", key)
	fmt.Println("Index of key: ", LinearSearchUnsortedData(numbers, key))
	fmt.Println("Index of 999: ", LinearSearchUnsortedData(numbers, 999))
	sort.Ints(numbers)
	fmt.Println("Sorted: ", numbers)
	fmt.Println("Index of key: ", LinearSearchSortedData(numbers, key))
	fmt.Println("Index of 999: ", LinearSearchSortedData(numbers, 999))
	fmt.Println("Index of key: ", BinarySearch(numbers, key))
	fmt.Println("Index of 999: ", BinarySearch(numbers, 999))
}

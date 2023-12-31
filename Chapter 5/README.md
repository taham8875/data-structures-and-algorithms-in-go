# Searching

Searching is the process of finding a particular item in a collection of items.

Computer sort very large sets of data, so we need efficient algorithms to search for items in these sets.

Storing the data is the proper state helps us to search for items efficiently. For example, search sorted data is much faster than searching unsorted data.

# Different Seaching Algorithms

- Linear Search - Unsorted Data
- Linear Search - Sorted Data
- Binary Search - Sorted Data
- String Search - Trie, Suffix Tree, Ternary Search Tree
- Hashing - Hash Table


# Linear Search - Unsorted Data

Linear search is the simplest search algorithm. It is also known as sequential search. We start at the beginning of the data structure and check each item until we find the item we are looking for or until we reach the end of the data structure.


# Linear Search - Sorted Data

If the data is sorted, we can stop searching as soon as we find an item that is greater than the item we are looking for. This is because all the items after that item will also be greater than the item we are looking for.

# Binary Search - Sorted Data

Binary search is a very efficient search algorithm. It is also known as half-interval search or logarithmic search.

In binary search, we start at the middle of the data structure and check if the item we are looking for is greater than, less than, or equal to the middle item. If it is equal, we have found the item we are looking for. If it is greater than, we search the right half of the data structure. If it is less than, we search the left half of the data structure. We repeat this process until we find the item we are looking for or until we reach the end of the data structure.

Code for linear search and binary search:

```go
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
```

> String Search and Hashing will be covered in a later chapter.


# How sorting is useful in Selection Algorithm?

Selecting and element from array is a O(1) operation no matter it is sorted or not, sorting the array may seem more expensive than selecting the element directly at first glance (O(n log n) in its best cases), but it is useful when you need to select multiple elements from the same array.

For example, imagine we want to select the maximum element, then the second maximum element, then the third maximum element, and so on. If we sort the array first, we can select the maximum element in O(1) time (get the last element), the second maximum element in O(1) time, and so on. If we don't sort the array first, we have to select the maximum element in O(n) time, the second maximum element in O(n) time, and so on.
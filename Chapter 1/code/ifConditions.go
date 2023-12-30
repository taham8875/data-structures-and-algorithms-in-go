package main

import "fmt"

func isGreaterThan(x, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}

func isAreaLessThanLimit(width, height, limit float32) bool {
	if area := width * height; area < limit {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isGreaterThan(10, 5))            // output: true
	fmt.Println(isAreaLessThanLimit(10, 5, 100)) // output: true
}

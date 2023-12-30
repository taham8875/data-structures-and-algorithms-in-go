// Given a list, you need to rotate its elements K number of times. For example, a list [10,20,30,40,50,60] rotate by 2 positions to [30,40,50,60,10,20]

package main

import "fmt"

func rotate(list []int, numberOfPositions int) {
	reverseArray(list, 0, numberOfPositions-1)
	reverseArray(list, numberOfPositions, len(list)-1)
	reverseArray(list, 0, len(list)-1)
}

func reverseArray(list []int, start int, end int) {
	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
}

func main() {
	list := []int{10, 20, 30, 40, 50}
	rotate(list, 2)
	fmt.Println(list)
}

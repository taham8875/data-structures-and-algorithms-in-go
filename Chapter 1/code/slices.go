package main

import "fmt"

func main() {
	var s []int
	fmt.Println("The value of s is", s) // output: The value of s is []

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	fmt.Println("The value of s is", s) // output: The value of s is [0 1 2 3 4 5 6 7 8 9]

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("The value of s2 is", s2) // output: The value of s2 is [0 1 2 3 4 5 6 7 8 9]

	s3 := make([]int, 10)
	fmt.Println("The value of s3 is", s3) // output: The value of s3 is [0 0 0 0 0 0 0 0 0 0]

	s4 := make([]int, 10, 20)                     // 10 is the length and 20 is the capacity
	fmt.Println("The value of s4 is", s4)         // output: The value of s4 is [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("The length of s4 is", len(s4))   // output: The length of s4 is 10
	fmt.Println("The capacity of s4 is", cap(s4)) // output: The capacity of s4 is 20

	s5 := s4[0:5]
	fmt.Println("The value of s5 is", s5)         // output: The value of s5 is [0 0 0 0 0]
	fmt.Println("The length of s5 is", len(s5))   // output: The length of s5 is 5
	fmt.Println("The capacity of s5 is", cap(s5)) // output: The capacity of s5 is 20
}

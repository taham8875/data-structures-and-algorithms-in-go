package main

import "fmt"

func main() {
	// Normal for loop
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice of integers

	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	fmt.Println("The sum of all numbers is", sum) // output: The sum of all numbers is 55

	// For loop which acts like a while loop

	i := 0
	sum = 0

	for i < len(numbers) {
		sum += numbers[i]
		i++ // otherwise we'll have an infinite loop
	}

	fmt.Println("The sum of all numbers is", sum) // output: The sum of all numbers is 55

	// For loop which acts like an infinite loop

	i = 0
	sum = 0

	for {
		sum += numbers[i]
		i++
		if i >= len(numbers) {
			break // otherwise we'll encounter an index out of range error
		}
	}

	fmt.Println("The sum of all numbers is", sum) // output: The sum of all numbers is 55

	// for range loop

	sum = 0

	for index, value := range numbers { // You can also use _ to ignore the index or value if you don't need it
		fmt.Println("The index is", index, "and the value is", value)
		sum += value
	}

	fmt.Println("The sum of all numbers is", sum) // output: The sum of all numbers is 55
}

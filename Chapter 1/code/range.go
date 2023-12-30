package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers {
		fmt.Println("The index is", index, "and the value is", value)
	}

	keyValuePairs := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cat",
	}

	for key, value := range keyValuePairs {
		fmt.Println("The key is", key, "and the value is", value)
	}

	s := "Hello World!"

	for index, character := range s {
		fmt.Println("The index is", index, "and the character is", string(character))
	}

}

package main

import "fmt"

func main() {
	fruitsPrices := map[string]int{
		"apple":  100,
		"banana": 200,
		"orange": 300,
	}

	for key, value := range fruitsPrices {
		fmt.Printf("The price of %s is %d\n", key, value)
	}

	fmt.Println("The price of an apple is", fruitsPrices["apple"])

	delete(fruitsPrices, "apple")

	fmt.Println("The price of an apple is", fruitsPrices["apple"])

	if value, ok := fruitsPrices["apple"]; ok {
		fmt.Println("The price of an apple is", value)
	} else {
		fmt.Println("Sorry, the apple is not available")
	}
}

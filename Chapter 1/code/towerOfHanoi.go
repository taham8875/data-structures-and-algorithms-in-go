package main

import "fmt"

func towerOfHanoi(numberOfDisks int, from string, to string, temp string) {
	if numberOfDisks == 1 {
		fmt.Println("Move disk 1 from rod", from, "to rod", to)
		return
	}
	towerOfHanoi(numberOfDisks-1, from, temp, to)
	fmt.Println("Move disk", numberOfDisks, "from rod", from, "to rod", to)
	towerOfHanoi(numberOfDisks-1, temp, to, from)
}

func main() {
	numberOfDisks := 3
	towerOfHanoi(numberOfDisks, "A", "C", "B")
}

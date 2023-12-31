package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	stack := stack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for stack.Len() != 0 {
		fmt.Println(stack.Pop(), " ") // Note the last in is the first out
	}
}

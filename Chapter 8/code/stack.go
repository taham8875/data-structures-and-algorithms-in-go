package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	stack []interface{}
}

func (stack *Stack) Push(value interface{}) {
	stack.stack = append(stack.stack, value)
}

func (stack *Stack) Pop() (value interface{}, err error) {
	if stack.isEmpty() {
		return nil, errors.New("empty stack")
	}
	value = stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return value, nil
}

func (stack *Stack) isEmpty() bool {
	return len(stack.stack) == 0
}

func (stack *Stack) Size() int {
	return len(stack.stack)
}

func (stack *Stack) Peek() interface{} {
	if stack.isEmpty() {
		return nil
	}
	return stack.stack[len(stack.stack)-1]
}

func main() {
	stack := new(Stack)
	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!!")
	fmt.Println(stack.Size())
	for stack.Size() > 0 {
		value, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(value)

	}

	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

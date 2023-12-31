package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
	queue := queue.New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	for queue.Len() != 0 {
		fmt.Println(queue.Dequeue(), " ") // Note the first in is the first out
	}
}

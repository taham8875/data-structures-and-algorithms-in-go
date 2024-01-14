package main

import (
	"errors"
	"fmt"
)

type List struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return list.count == 0
}

func (list *List) AddHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *List) AddTail(value int) {
	newNode := &Node{value, nil}
	current := list.head

	if current == nil {
		list.head = newNode
		list.count++
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	list.count++
}

func (list *List) Print() {
	for node := list.head; node != nil; node = node.next {
		fmt.Println(node.value)
	}
}

func (list *List) IsPresent(value int) bool {
	for temp := list.head; temp != nil; temp = temp.next {
		if temp.value == value {
			return true
		}
	}
	return false
}

func (list *List) DeleteHead() (int, error) {
	if list.IsEmpty() {
		return 0, errors.New("empty list")
	}

	value := list.head.value
	list.head = list.head.next
	list.count--
	return value, nil
}

func (list *List) DeleteNode(value int) bool {
	if list.IsEmpty() {
		return false
	}

	if list.head.value == value {
		list.head = list.head.next
		list.count--
		return true
	}

	for temp := list.head.next; temp != nil; temp = temp.next {
		if temp.next.value == value {
			temp.next = temp.next.next
			list.count--
			return true
		}
	}
	return false
}

func (list *List) DeleteNodes(value int) {
	if list.head == nil {
		return
	}

	for temp := list.head; temp.next != nil; {
		if temp.next.value == value {
			temp.next = temp.next.next
			list.count--
		} else {
			temp = temp.next
		}
	}
}

func (list *List) DeleteList() {
	list.head = nil
	list.count = 0
	// garbage collector will remove the entire list
}

func (list *List) Reverse() {
	current := list.head
	var previous, next *Node
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	list.head = previous
}

func (list *List) RemoveDuplicates() {
	current := list.head
	for current != nil {
		if current.next != nil && current.value == current.next.value {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

func (list *List) Copy() *List {
	var headNode *Node

	copyList := new(List)

	current := list.head

	if current == nil {
		copyList.head = nil
		return copyList
	}

	headNode = &Node{current.value, nil}

	tempNode := headNode

	for current != nil {
		tempNode.value = current.value
		if current.next != nil {
			tempNode.next = &Node{0, nil}
		}
		current = current.next
		tempNode = tempNode.next
	}

	copyList.head = headNode

	return copyList
}

func (list1 *List) Compare(list2 *List) bool {
	current1 := list1.head
	current2 := list2.head

	for {
		if current1 == nil && current2 == nil {
			return true
		}
		if current1 == nil || current2 == nil || current1.value != current2.value {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}
}

func (list *List) Length() int {
	return list.count
}

func (list *List) at(index int) (int, error) {
	if index < 0 || index > list.Length() {
		return 0, errors.New("index out of bound")
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func main() {
	list := &List{} // or list := new(List)
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)
	list.AddHead(4)
	list.AddHead(5)
	list.Print()
	list2 := list.Copy()
	list2.AddHead(6)
	list2.Print()
	fmt.Println(list.Compare(list2))
	fmt.Println(list.Length())
	if value, err := list.at(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

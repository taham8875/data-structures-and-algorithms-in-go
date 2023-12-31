# Abstract data type & Go Collections

# Abstract data type

An abstract data type (ADT) is a logical description of hwo we view the data and the operations that are allowed without regard to how they will be implemented. This means that we are concerned only with what the data is representing and not with the actual implementation.

In other words, an ADT is defined as the users point of view.

# Data Structure

Data structure is a specialized format for organizing, processing, retrieving and storing data. It represents how the data is stored in the memory.

All data structures have their own advantages and disadvantages. It is important to know the characteristics of the data structure so that we can use the appropriate data structure for the problem we are trying to solve.

examples:

- Array
- Linked List
- Stack
- Queue
- Tree
- Graph
- Hash Table

# Go Collections Framework

Go programming language provides a rich set of collections, which are high quality, high performant and reusable data structures and algorithms.

The advantages of using the go collections framework are:

- You don't have to write your own data structures and algorithms repeatedly. Don't reinvent the wheel.
- The collections are well tested and optimized for performance.
- Development time is reduced.
- Easy to read and review the code since it is well documented.

# Stack

A stack is a linear data structure that follows the LIFO (Last-In-First-Out) principle. This means that the last element inserted inside the stack is removed first.

Application of stack:

- Recursion: The function call is stored in system stack.
- Postfix evaluation

Common operations on stack:

- `push(element)` - Insert an element at the top of the stack.
- `pop()` - Remove an element from the top of the stack.
- `top()` - Return the top element of the stack.
- `size()` - Return the number of elements in the stack.
- `isEmpty()` - Return true if the stack is empty.

> Note that all these operations are O(1) time complexity.

Stack usage from go collections framework:

```go
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

```

# Queue

A queue is a linear data structure that follows the FIFO (First-In-First-Out) principle. This means that the first element inserted inside the queue is removed first.

Common operations on queue:

- `add(element)` - Insert an element at the end of the queue.
- `remove()` - Remove an element from the front of the queue.
- `front()` - Return the front element of the queue.
- `size()` - Return the number of elements in the queue.
- `isEmpty()` - Return true if the queue is empty.

> Note that all these operations are O(1) time complexity.

Queue usage from go collections framework:

```go
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

```

# Tree

Tree is a hierarchial data structure. The top element of the tree is called the root, every element except the root has exactly one parent, and zero or more child elements

## Binary Tree

A binary tree is a tree in which each node has at most two children, which are referred to as the left child and the right child.

## Binary Search Tree

A binary search tree is a binary tree in which:

- The value of the left child is less than the value of its parent node.
- The value of the right child is greater than or equal to the value of its parent node.

Common operations on binary search tree:

- `insert(element)` - Insert an element into the tree.
- `delete(element)` - Delete an element from the tree.
- `search(element)` - Search for an element in the tree.
- `findMin()` - Return the minimum element in the tree.
- `findMax()` - Return the maximum element in the tree.

> Note: when the tree is balanced, the time complexity of all these operations are O(log n). The worst case is O(n) when the tree is skewed.

A binary search tree is skewed when the tree is not balanced:

- Right skewed: all the nodes have either one or no right child.
- Left skewed: all the nodes have either one or no left child.

# Priority Queue (Heap)

Heap is a special tree-based data structure in which the tree is a complete binary tree. A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.

There are two types of heaps:

- Max heap: Each node is greater than or equal to its children.
- Min heap: Each node is less than or equal to its children.

You may thing that we something like that while implementing the heap

```go
type Node struct {
    value int
    left  *Node
    right *Node
}
```

But it is much simpler, heap can be implemented using an array.

Heap is useful when you want to get the maximum or minimum element.

Common operations on heap:

- `insert(element)` - Insert an element into the heap.
- `delete(element)` - Delete an element from the heap.
- `heapifyUp()` - Heapify the heap from bottom to top.
- `heapifyDown()` - Heapify the heap from top to bottom.

# Hash Table

Hash table is a data structure that maps keys to values. It uses a hash function to compute an index os a list.

The process of storing objects in a hash table:

1. Create a list of size N to store the objects. The list is called a hash table.
2. Find the hash code of any object by passing it through a hash function.
3. Take the modulus of the hash code with the size of the list. This will give you the index of the list where the object should be stored.
4. Finally store the object at the computed index.

Common operations on hash table:

- `Insert(key, value)` - Insert a key-value pair into the hash table.
- `Delete(key)` - Delete a key-value pair from the hash table.
- `Search(key)` - Search for a key in the hash table.

> Note: The time complexity of all these operations is O(1).

Hash tables are built into the Go language. You can use the built-in map type to create a hash table.

```go
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

```

Note that order is not guaranteed when iterating over a map.

# Set

Set is a data structure that stores unique elements. It does not allow duplicate elements.

In go collections set is implemented on the top of map.

example:

```go
package main

import (
	"fmt"

	"github.com/golang-collections/collections/set"
)

func main() {
	set := set.New()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(3) // Duplicate entry

	fmt.Println(set.Has(1)) // true
	fmt.Println(set.Has(3)) // true
	fmt.Println(set.Has(5)) // false

}
```

# Dictionary / Symbol Table

Dictionary is a data structure that stores key-value pairs.

# Binary Search Tree For Strings

The simplest way to implement a symbol table is to use a binary search tree. The keys are stored in sorted order in the tree. The left subtree of a node contains keys that are less than the node's key, and the right subtree of a node contains keys that are greater than the node's key.

If the key is balanced, the lookup is logarithmic.

# Hash Table

Hash table is a data structure that maps keys to values. It uses a hash function to compute an index of a list.

Hash table is excellent for constant time lookup.

Now let's imagine we want to implement a autocomplete feature like in google search:

- BST are not a good solution since the tree does not know if it should traverse the right or the left tree
- Hash table is not a good solution since passing part of the string to the hash function will not give us the same index

Trie and Ternary search tree are good solutions for this problem.

# Trie

Tree is like a tree where we store only one character at each node and store the final key value pair at the leaves.

Trie is an efficient data structure since we can search for a key in O(L) time where L is the maximum length of the string.

# Ternary Search Tree / Ternary Search Trie

Although trie have a performant lookup, it is not space efficient since each node has 26 pointers. Ternary search tree (TST) is a space efficient trie.

In TST each node contains a character, three pointers and a value. The three pointers are for the current char hold by the node (equal), chars less than and chars greater than the current char.

# Graph

Graph is a data structure that represents a network consists of a finite set of vertices (or nodes) and a set of edges connecting these vertices.

Edges can either be directed or undirected. If the edges are directed you can move in one direction only. If the edges are undirected you can move in both directions.

# Depth First Search (DFS)

In depth first search (DFS) we start from a vertex and then visit all the adjacent vertices of the current vertex until we reach a dead end. Once we reach a dead end, we backtrack to the last vertex that has an adjacent vertex and continue this process until all the vertices are visited.

We either get the next vertex from a stack or from a recursive call stack.

# Breadth First Search (BFS)

In BFS the graph is traversed in layer be layer fashion. We start from a vertex and visit all the adjacent vertices of the current vertex. Once we visit all the adjacent vertices of the current vertex, we move to the next vertex from the queue and continue this process until all the vertices are visited.

# Sorting Algorithms

Sorting is the process of arranging the elements of a collection in a certain order.

There are efficient sorting algorithms that can sort a collection in O(n log n) time. So when we call a `sort()` method on a collection, we expect the collection to be sorted in O(n log n) time.

# Counting Sort

Counting sort is a sorting algorithm that sorts the elements of an array by counting the number of occurrences of each unique element in the array. The count is stored in an auxiliary array and the sorting is done by mapping the count as an index of the auxiliary array.

Counting sort is used when we know in advance the range of the elements in the array. (e.g. the ages for group of people, which can be estimated to be between 0 and 150)

Counting sort runs in O(n) time complexity.

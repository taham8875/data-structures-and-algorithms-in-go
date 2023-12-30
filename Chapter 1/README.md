# Chapter 1 - Introduction - Programming Overview

# Hello World

Your first go program

```go
package main

import "fmt"

// This is a single line comment

/*
This is a multi-line comment
You can have as many lines as you want
*/

func main() {
	fmt.Println("Hello World!")
}

```

1. The go programming language uses packages to organize the associated source files together. Note that the `main()` function is in the `main` package.
2. Go uses the `import` keyword to include third-party packages. In this case, we are importing the `fmt` package, which contains functions for printing to the console.
3. The `main()` function is the entry point of the program. When the program is executed, the control is transferred to the `main()` function.
4. Comments are ignored by the compiler. `//` is used for single line comments and `/* */` is used for multi-line comments.

# Variables and Constants

Variables are storage locations that holds data of a specific type. When variable is declared, some memory is allocated for it. The size of that memory depends on the type of the variable.

Variables are declared using the `var` keyword. The syntax is:

```go
var <variable_name> <type>
```

Note that go is statically typed language, which means that the type of the variable is known at compile time. The type of the variable cannot be changed later.

Example:

```go
package main

import "fmt"

func main() {
	var variable1 int
	var variable2 int
	variable1 = 100
	fmt.Println("The value of variable1 is", variable1) // output: The value of variable1 is 100
	fmt.Println("The value of variable2 is", variable2) // output: The value of variable2 is 0
}

```

The are no uninitialized variables in go. If a variable is not initialized, it is assigned a zero value.

The zero value depends on the type of the variable. For example:

- the zero value of an integer is 0
- the zero value of a floating point number is 0.0
- the zero value of a boolean is false
- the zero value of a string is ""
- the zero value of a pointer is `nil`

Constants are declared using the `const` keyword. The syntax is:

```go
const <constant_name> <type> = <value>
```

`:=` operator allows short declaration of variables. The type of the variable is inferred from the value on the right hand side of the `:=` operator.

# Basic Data Types

1. **Boolean** - `bool` - true or false
2. **Integer** - `int` - 32 or 64 bits (other variants are `int8`, `int16`, `int32`, `int64`)
3. **Unsigned Integer** - `uint` - 32 or 64 bits (other variants are `uint8`, `uint16`, `uint32`, `uint64`)
4. **Floating Point** - `float32` or `float64`
5. **Complex** - `complex64` or `complex128`
6. **String** - `string` - sequence of unicode characters
7. **Byte** - `byte` - alias for `uint8`
8. **Rune** - `rune` - alias for `int32` and used to store unicode characters

# Strings

Strings are sequence of unicode characters.

Double quotes are used to create strings. Example:

```go
s := "Hello World!"
```

Strings are immutable in go. Once a string is created, it cannot be changed.

To change a string, you need to convert it to a slice of runes, which is a mutable sequence of unicode characters. then convert the slice of runes back to a string.

```go
package main

import "fmt"

func main() {
	s := "hello World!"
	// s[0] = "H" // error: cannot assign to s[0] (strings are immutable)
	fmt.Println(s) // output: hello World!
	r := []rune(s)
	fmt.Println(r) // output: [104 101 108 108 111 32 87 111 114 108 100 33] // each rune is a Unicode code point
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2) // output: Hello World!
}

```

There are many useful operations that can be performed on strings. Some of them are:

- `len(s)` - returns the length of the string
- `s[i]` - returns the ith character of the string
- `s + t` - concatenates two strings (e.g. `s := "Hello"`, `t := "World!"`, `s + " " + t = "Hello World!"`)
- `"hello" == "hello" // true` - compares two strings
- `"a" < "b" // true` - compares two strings lexicographically
- `s[i:j]` - returns a substring from index `i` to `j-1` (slicing)

# Conditional Statements

## If Statement

The syntax of the if statement is:

```go
if <condition> {
    // code to be executed if condition is true
}
```

If statement can have an optional `else` block. The syntax is:

```go
if <condition> {
    // code to be executed if condition is true
} else {
    // code to be executed if condition is false
}
```

example:

```go
func isGreaterThan(x, y int) bool {
    if (x > y) {
        return true
    } else {
        return false
    }
}
```

> A more short and concise way to write the above function is:
> `return x > y`

Go if condition supports a special syntax called `if with a short statement`. The syntax is:

```go
if <short statement>; <condition> {
    // code to be executed if condition is true
}
```

example:

```go
func isAreaLessThanLimit(width, height, limit float32) bool {
	if area := width * height; area < limit {
		return true
	} else {
		return false
	}
}
```

## Switch Statement

The syntax of the switch statement is:

1. switch with basic syntax

```go
switch <optional initialization>; <expression> {
    case <expression1>:
        // code to be executed if expression == expression1
    case <expression2>:
        // code to be executed if expression == expression2
    default:
        // code to be executed if expression is different from all expressions
}
```

example:

```go
package main

import "fmt"

func main() {
	switch i := 1; i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
    case 4, 5, 6, 7, 8, 9: // multiple cases
        fmt.Println("Four through Nine")
	default:
		fmt.Println("Unknown Number")
	}
}

```

2. switch with condition

```go
switch {
    case <condition1>:
        // code to be executed if condition1 is true
    case <condition2>:
        // code to be executed if condition2 is true
    default:
        // code to be executed if all conditions are false
}
```

For example, lets solve the fizzbuzz problem using switch statement.

> Take in a number from 1 through 1000, print Fizz if the number is divisible by 3, print Buzz if the number is divisible by 5, or FizzBuzz if the number is divisible by both. If none of these conditions are true, print nothing.

```go
package main

import "fmt"

// Take in a number from 1 through 1000, print Fizz if the number is divisible by 3, print Buzz if the number is divisible by 5, or FizzBuzz if the number is divisible by both. If none of these conditions are true, print nothing.

func main() {
	var i int
	fmt.Scanf("%d\n", &i)

	switch {
	case i%3 == 0 && i%5 == 0:
		fmt.Println("FizzBuzz")
	case i%3 == 0:
		fmt.Println("Fizz")
	case i%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println("")
	}

}

```

3. switch with type

```go
switch <var>.(type) {
    case <type1>:
        // code to be executed if expression is of type type1
    case <type2>:
        // code to be executed if expression is of type type2
    default:
        // code to be executed if expression is different from all types
}
```

example:

```go
	var s interface{} = "Hello"

	switch s.(type) {
	case string:
		fmt.Println("s is a string")
	case int:
		fmt.Println("s is an int")
	default:
		fmt.Println("Unknown type")
	}
```

> `interface{}` is a special type in go that can hold values of any type. If you familiar with typescript, it is similar to `any` type.

# For Loop

There are four forms of for loop in go.

1. `for <initialization>; <condition>; <increment/decrement>` {}
2. `for <condition> {}` - similar to while loop
3. `for {}` - infinite loop
4. `for <index>, <value> := range <collection>` {} - iterate over a collection

```go
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

```

## range

The `range` keyword is used to iterate over a collection. The syntax is:

```go
for <index>, <value> := range <collection> {
    // code to be executed
}
```

You can omit the index or value if you don't need it. For example:

```go
for _, value := range numbers {
    // code to be executed
}
```

Behavior of `range` differs based on the type of the collection.

1. with arrays and slices, `range` returns the index and value of each element in the collection
2. with maps, `range` returns the key and value pairs of each element in the collection
3. with strings, `range` returns the index and unicode code point of each character in the string

# Functions

Functions are used ti provide modularity to the program. Functions are declared using the `func` keyword. The syntax is:

```go
func <function_name>(<parameter_list>) <return_type> {
    // code to be executed
}
```

example:

```go
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

# Pointers

Pointers are used to store the memory address of a variable. The type of the pointer is the type of the variable it points to.

Pointers are used to access the value of the variable they point to. This is called dereferencing the pointer.

example:

```go
package main

import "fmt"

func main() {
	data := 20
	ptr := &data

	fmt.Println("Value stored at variable data = ", data)
	fmt.Println("Address of variable data = ", &data)
	fmt.Println("Value stored at variable ptr = ", ptr)
	fmt.Println("Value stored at variable *ptr = ", *ptr)

}
```

# Parameter Passing

## Call by Value

Arguments can be passed to the function using parameters, by default, all parameters are passed by value.

This means that a separate copy of the variable is created and passed to the function. Any changes made to the parameter inside the function does **not** affect the original variable.

example:

```go
package main

import "fmt"

func incrementPassByValue(x int) {
	x++
}

func main() {
	x := 0
	fmt.Println("x before calling incrementPassByValue:", x) // output: x before calling incrementPassByValue: 0
	incrementPassByValue(x)
	fmt.Println("x after calling incrementPassByValue:", x)      // output: x after calling incrementPassByValue: 0
}
```

## Call by Reference

Arguments can be passed to the function using pointers, this is called call by reference.

This means that the address of the variable is passed to the function. Any changes made to the parameter inside the function affects the original variable.

example:

```go
package main

import "fmt"

func incrementPassByReference(x *int) {
	*x++
}

func main() {
	x := 0
	fmt.Println("x before calling incrementPassByReference:", x) // output: x before calling incrementPassByReference: 0
	incrementPassByReference(&x)
	fmt.Println("x after calling incrementPassByReference:", x) // output: x after calling incrementPassByReference: 1
}
```

# Structures

Structures are used to group related data together. Structures are declared using the `struct` keyword. The syntax is:

```go
type <struct_name> struct {
    <field_name> <field_type>
    <field_name> <field_type>
    ...
}
```

example:

```go
package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	newStudent := student{1, "John"}
	fmt.Println("newStudent:", newStudent) // output: newStudent: {1 John}
	// You can access inner fields of a struct using the dot operator
	fmt.Println("newStudent.name:", newStudent.name) // output: newStudent.name: John
	// If you had a pointer to a struct, you could access inner fields using the dot operator as well
	// Go automatically dereferences pointers when using the dot operator
	newStudentPointer := &newStudent
	fmt.Println("newStudentPointer.name:", newStudentPointer.name) // output: newStudentPointer.name: John

	// You can also create a struct using the new keyword which returns a pointer to the struct
	newStudent2 := new(student)
	// Initialize the fields of the struct with the default values
	fmt.Println("newStudent2:", newStudent2) // output: newStudent2: &{0 }
	// Now you can set the fields of the struct
	newStudent2.id = 2
	newStudent2.name = "Jane"
	fmt.Println("newStudent2:", newStudent2) // output: newStudent2: &{2 Jane}

	// -------

	fmt.Println(student{id: 3, name: "Jack"}) // Named initialization // output: {3 Jack}
	fmt.Println(student{name: "Jack", id: 3}) // Order doesn't matter // output: {3 Jack}
	fmt.Println(student{name: "Jack"})        // output: {0 Jack} // Note the default value for id
}
```

# Methods

Methods are functions that are associated with a struct. Methods are declared using the `func` keyword followed by the receiver, we can access the method associated with a struct using the dot operator `.`

example:

```go
package main

import "fmt"

type Rectangle struct {
	width  float32
	height float32
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.width * rectangle.height
}

func (rectangle Rectangle) Perimeter() float32 {
	return 2 * (rectangle.width + rectangle.height)
}

func main() {
	r := Rectangle{5.0, 4.0}
	fmt.Println("Rectangle is: ", r)
	fmt.Println("Rectangle's area is: ", r.Area())
	fmt.Println("Rectangle's perimeter is: ", r.Perimeter())

	rp := &r
	fmt.Println("Rectangle is: ", rp)
	fmt.Println("Rectangle's area is: ", rp.Area())
	fmt.Println("Rectangle's perimeter is: ", rp.Perimeter())
}

```

There are two ways of defining associated methods with a struct.

1. Accessor functions (pass by value)

```go
func (rectangle Rectangle) someFunction() float32 {
	// code to be executed
}
```

2. Mutator (modifier) functions (pass by reference)

```go
func (rectangle *Rectangle) someFunction() float32 {
    // code to be executed
}
```

# Interfaces

Interfaces are used to define behavior of an object. Interfaces are declared using the `interface` keyword, it defines a set of methods with its return type.

Syntax:

```go
type <interface_name> interface {
    <method_name1>(<parameter_list>) <return_type>
    <method_name2>(<parameter_list>) <return_type>
    ...
}
```

To implement an interface, a type must implement all the methods defined in the interface.

example:

```go
package main

import "fmt"

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Circle struct {
	radius float32
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.width * rectangle.height
}

func (rectangle Rectangle) Perimeter() float32 {
	return 2 * (rectangle.width + rectangle.height)
}

func (circle Circle) Area() float32 {
	return 3.14 * circle.radius * circle.radius
}

func (circle Circle) Perimeter() float32 {
	return 2 * 3.14 * circle.radius
}

// Because both Rectangle and Circle implement the Shape interface, we can use them interchangeably.

func TotalArea(shapes ...Shape) float32 {
	var area float32
	for _, shape := range shapes {
		area += shape.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float32 {
	var perimeter float32
	for _, shape := range shapes {
		perimeter += shape.Perimeter()
	}
	return perimeter
}

func main() {
	rectangle := Rectangle{5.0, 4.0}
	circle := Circle{10}

	fmt.Println("Rectangle is: ", rectangle)
	fmt.Println("Rectangle's area is: ", rectangle.Area())
	fmt.Println("Rectangle's perimeter is: ", rectangle.Perimeter())

	fmt.Println("Circle is: ", circle)
	fmt.Println("Circle's area is: ", circle.Area())
	fmt.Println("Circle's perimeter is: ", circle.Perimeter())

	fmt.Println("Total area of shapes is", TotalArea(rectangle, circle))
	fmt.Println("Total perimeter of shapes is", TotalPerimeter(rectangle, circle))
}
```

# Arrays

Arrays are used to store a fixed-size sequential collection of elements of the same type. Arrays are declared using the `var` keyword. The syntax is:

```go
var <array_name> [<size>] <type>
```

example:

```go
package main

import "fmt"

func main() {
	var array [10]int
	fmt.Println("The value of array is", array) // output: The value of array is [0 0 0 0 0 0 0 0 0 0]
	for i, _ := range array {
		array[i] = i
	}
	fmt.Println("The value of array is", array)       // output: The value of array is [0 1 2 3 4 5 6 7 8 9]
	fmt.Println("The length of array is", len(array)) // output: The length of array is 10
}
```

Note that array variable is the whole array not just a pointer to its first element like in C/C++.

# Slices

Arrays are fixed in size and don't support dynamic resizing. Slices are used to overcome this limitation. Slices are abstraction over arrays. i.e. it uses arrays as its underlying data structure.

The common operations on slices are:

1. `append()` - adds an element to the a slice, if the size of the underlying array is not enough, a new array is created and the elements are copied to the new array.
2. `len()` - returns the number of elements in the slice
3. `cap()` - returns the capacity of the underlying array
4. `copy()` - copies elements from one slice to another
5. `s[start:end]` - returns a slice of the slice `s` from index `start` to `end-1` (re-slice)

There are two ways to create a slice:

1. Using the `make()` function

```go
slice := make([]<type>, <size>, <capacity>)
```

2. Declare it as an array without specifying the size

```go
var slice []<type>
```

# Map / Dictionary

A map is a collection of key-value pairs. Maps are declared using the `map` keyword. The syntax is:

```go
var <map_name> map[<key_type>]<value_type>
```

map has to be initialized before it can be used using `make`. The syntax is:

```go
<map_name> := make(map[<key_type>]<value_type>)
```

Common operations on maps are:

1. Assigning a value to a key - `map[key] = value`
2. Deleting a key-value pair - `delete(map, key)`
3. Accessing a value using a key - `value, ok := map[key]` - if the key exists, `ok` is true, otherwise `ok` is false

Note that if the key doesn't exist, the value returned is the zero value of the value type, so `value, ok` is used to check if the key exists and its value
coincidentally is the zero value of the value type.

Note that the order in which the key-value pairs are stored in the map is not guaranteed.

example:

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

# Stack

Stack is a data structure that follows the LIFO (Last In First Out) principle. `pop()` and `push()` are the two main operations on a stack.

# System Stack and Method Calls

When a method is called, the current execution is stopped and the control goes ti the called method. The current execution is saved in the system stack. When the called method returns, the execution is resumed from the point where it was stopped.

# Recursive Functions

A recursive function is function that calls itself

It consists of two parts:

1. Termination condition - the condition that stops the recursion, in which process a simple case and does not call itself
2. Body - the part of the function that calls itself, which include the recursive expansion

Three important properties of a recursive function:

1. It must have a termination condition (without it, the function will call itself infinitely)
2. It must change its state and move towards the termination condition
3. It must call itself.

> Note: the speed of the recursive function is slower than the iterative function because of the stack overhead. It the problem can be solved using iterative solution (using loops), it is better to use it.

example (factorial):

```go
package main

import "fmt"

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
    fmt.Println("The factorial of 5 is", factorial(5)) // output: "The factorial of 5 is 120"
}

```

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

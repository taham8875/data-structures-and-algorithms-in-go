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

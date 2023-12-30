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

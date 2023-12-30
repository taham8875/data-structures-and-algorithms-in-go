package main

import "fmt"

func incrementPassByValue(x int) {
	x++
}

func incrementPassByReference(x *int) {
	*x++
}

func main() {
	x := 0
	fmt.Println("x before calling incrementPassByValue:", x) // output: x before calling incrementPassByValue: 0
	incrementPassByValue(x)
	fmt.Println("x after calling incrementPassByValue:", x)      // output: x after calling incrementPassByValue: 0
	fmt.Println("x before calling incrementPassByReference:", x) // output: x before calling incrementPassByReference: 0
	incrementPassByReference(&x)
	fmt.Println("x after calling incrementPassByReference:", x) // output: x after calling incrementPassByReference: 1
}

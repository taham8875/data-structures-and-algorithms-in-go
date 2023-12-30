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
	fmt.Println(s2)                                              // output: Hello World!
	fmt.Println("The length of s is", len(s))                    // output: The length of s is 12
	fmt.Println("The fifth character of s is", s[4])             // output: The fifth character of s is 111
	fmt.Println("The fifth character of s is", string(s[4]))     // output: The fifth character of s is o
	fmt.Println("Concatenating s and s2:", s+s2)                 // output: Concatenating s and s2: hello World!Hello World!
	fmt.Println("Is s equal to s2? ", s == s2)                   // output: Is s equal to s2? false
	fmt.Println("Is s equal to s", s == s)                       // output: Is s equal to s true
	fmt.Println("Is 'a' greater than 'b'? ", 'a' > 'b')          // output: Is 'a' greater than 'b'?  false
	fmt.Println("Slice the first five characters from s", s[:5]) // output: Slice the first five characters from s hello
}

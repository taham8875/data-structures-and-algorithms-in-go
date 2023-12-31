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

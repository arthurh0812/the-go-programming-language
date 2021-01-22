// the-new-function experiments with the built-in function "new"
package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 8
	fmt.Println(*p)

	q := newInt()
	fmt.Println(*q)
	*q = 5
	fmt.Println(*q)

	y := new(string)
	z := new(string)
	fmt.Println(y == z)
}

func newInt() *int {
	var dummy int
	return &dummy
}

func delta(old, new int) int {
	return new - old
	// q := new(string) doesn't work
}

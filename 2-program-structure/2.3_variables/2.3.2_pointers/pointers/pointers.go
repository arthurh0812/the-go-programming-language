package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x         // p of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	var p2 *int = &x
	fmt.Println(p2)

	if p == p2 {
		fmt.Println("EQUALITY.")
	}

	var p3 = f()
	fmt.Println(p3)
	fmt.Println(f() == f()) // "false"

	v := 1
	incr(&v)
	fmt.Println(incr(&v))
	fmt.Println(v)

	y := 2
	badIncr(y)
	fmt.Println(badIncr(y + 1))
	fmt.Println("Y is:", y)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func badIncr(p int) int {
	for i := 0; i < p; i++ {
		fmt.Println("I am coming from badIncr()!")
	}
	return p
}

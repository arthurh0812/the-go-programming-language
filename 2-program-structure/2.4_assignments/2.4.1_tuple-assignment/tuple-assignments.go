package main

import (
	"fmt"
)

func main() {
	var x, y = 6, 9

	x, y = y, x // y and x on the right-hand side are evaluated first

	d := gcd(18, 9)
	fmt.Println(d)

	f := fib(60)
	fmt.Println(f)

	typeTester(person{name: "James Bond", age: 32})

	typeTester(56.0)

	typeTester("I like Go!")

	typeTester(float32(90))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

type person struct {
	name string
	age  int
}

func typeTester(t interface{}) {
	if p, ok := t.(person); ok {
		fmt.Printf("%v: Type person\n", p)
		return
	}

	if i, ok := t.(int); ok {
		fmt.Printf("%v: Type int\n", i)
		return
	}

	if i8, ok := t.(int8); ok {
		fmt.Printf("%v: Type int8\n", i8)
		return
	}

	if i16, ok := t.(int16); ok {
		fmt.Printf("%v: Type int16\n", i16)
		return
	}

	if i32, ok := t.(int32); ok {
		fmt.Printf("%v: Type int32\n", i32)
		return
	}

	if i64, ok := t.(int); ok {
		fmt.Printf("%v: Type int64\n", i64)
		return
	}

	if s, ok := t.(string); ok {
		fmt.Printf("%q: Type string\n", s)
		return
	}

	if fl32, ok := t.(float32); ok {
		fmt.Printf("%v: Type float32\n", fl32)
		return
	}

	if fl64, ok := t.(float64); ok {
		fmt.Printf("%v: Type float64\n", fl64)
		return
	}

	fmt.Println("Sorry, we couldn't recognize this data type.")
}

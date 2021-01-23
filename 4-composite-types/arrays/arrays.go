package main

import "fmt"

var a [3]int // array of 3 intergers

func main() {
	fmt.Println(a[0])        // print the first element (index 0)
	fmt.Println(a[len(a)-1]) // print the last element (index (length-1)), a[2]

	// print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q, r)
	fmt.Println(r[2]) // "0" -> zero value of type int

	q = [...]int{2, 3, 4}
	// q = [4]int{1, 2, 3, 4} // Compile Error: cannot assign [4]int to [3]int

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[USD])

	hundred := [...]int{99: -1}

	fmt.Println(hundred)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 3}
	fmt.Println(a == b, a == c, b == c) // "true false true"
	d := [3]int{1, 2}
	fmt.Println(d)
	// fmt.Println(a == d) // compile error: cannot compare [2]int with [3]int

	var array = [32]byte{1, 4, 5, 7, 3, 10, 2, 2, 34, 53, 1, 26, 73, 24, 23, 43, 88, 26, 75, 82}
	fmt.Println(array)
	zero(&array)
	fmt.Println(array)
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

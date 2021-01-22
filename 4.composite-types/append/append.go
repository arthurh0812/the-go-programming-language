package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is a room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling (except for 0), for amortized linear complexity.
		zcap := zlen
		if len(x) > 0 {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...) // append all the elements of a slice z
	fmt.Println(z)

	var a, b []int
	for i := 0; i < 10; i++ {
		b = improvedAppendInt(a, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(b), b)
		a = b
	}
}

func improvedAppendInt(x []int, y ...int) []int {
	var z []int

	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// enough space in the underlying array
		z = x[:zlen]
	} else {
		// allocate new underlying array with at least added size (and double-size if smaller than the double length)
		zcap := zlen

		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

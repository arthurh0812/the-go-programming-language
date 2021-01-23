package main

import "fmt"

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	reverse(a[:])

	fmt.Println(a)

	s := []int{1, 2, 3, 4, 5, 6}

	// shift elements by using reverse()
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)

	fmt.Println(s)

	b := equal([]string{"cats", "dog", "rabbits"}, []string{"cats", "dogs", "rabbits"})

	fmt.Println(b)

	if s == nil {
		fmt.Println("only possible comparison with slices")
	}
}

// equal tests whether two slices contain the same elements
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings. The underlying array is modified during the function execution.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	myList := []string{"hello", "ok", "cool", "nice", "", "stop", "", ""}

	fmt.Println(myList)

	nonempty(myList)

	fmt.Println("1:", myList)

	fmt.Println("2:", nonempty2(myList))
}

// nonempty returns a slice holding only the non-empty strings. The underlying array is not modified during the function execution.
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

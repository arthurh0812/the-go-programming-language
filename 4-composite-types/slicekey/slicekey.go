package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

// Add increases the number of its list key of the map by one.
func Add(m map[string]int, list []string) {
	m[k(list)]++
}

// Count returns the number of added lists to a map.
func Count(m map[string]int, list []string) int {
	return m[k(list)]
}

func main() {
	myList1 := []string{"hello", "bye", "nice", "perfect"}
	myList2 := []string{"nothing"}

	Add(m, myList1)
	Add(m, myList2)
	Add(m, myList1)

	fmt.Println(m)

	n := Count(m, myList2)
	fmt.Println(n)

	n = Count(m, myList1)
	fmt.Println(n)
}

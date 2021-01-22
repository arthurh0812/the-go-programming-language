package main

import (
	"fmt"
	"sort"
)

// map type
var a = map[string]int{}

// nil map reference
var zeroMap map[string]int

func main() {
	ages := make(map[string]int) // mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34

	// equivalent to

	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	ages["alice"] = 32         // write to
	fmt.Println(ages["alice"]) // read from
	delete(ages, "alice")      // delete from

	fmt.Println(ages["alice"]) // 0

	ages["bob"] = ages["bob"] + 1 // happy first birthday!
	fmt.Println(ages["bob"])      // 1
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages["bob"]) // 3

	// _ = &ages["bob"] // compile error: cannot take address of map element

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	sortAges(ages)

	fmt.Println(zeroMap == nil)    // true
	fmt.Println(len(zeroMap) == 0) // true

	// lookups, delete, length and range are safe to use on nil map reference, but storing to nil map ref causes panic:
	zeroMap["daniel"] = 32 // panic: assignment to entry in nil map

	age, ok := ages["bob"]
	if !ok {
		// "bob" is not a key in this map, age = 0
	} else {
		fmt.Println(age)
	}

	if age, ok := ages["bob"]; ok {
		// only if there was key "bob" in map
		fmt.Println(age)
	}
}

func sortAges(ages map[string]int) {
	var names = make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, av := range a {
		if bv, ok := b[k]; !ok || av != bv {
			return false
		}
	}
	return true
}

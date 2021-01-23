package main

import "fmt"

func removeAdjacent(strings []string) []string {
	var z []string
	var lastString string
	for i, v := range strings {
		if strings[i] == lastString {
			continue
		}
		z = append(z, v)
		lastString = v
	}
	return z
}

func main() {
	str := []string{"ok", "hello", "hello", "hello", "ok", "nothin", "right", "right", "nothin"}

	dedup := removeAdjacent(str)

	fmt.Println(dedup)
}

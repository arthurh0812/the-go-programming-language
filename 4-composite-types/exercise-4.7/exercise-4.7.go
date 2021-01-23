package main

import "fmt"

func reverse(bs []byte) []byte {
	// get the unicode code points from bs
	runes := []rune(string(bs))

	// reverse the unicode code points
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// put them back into bs
	return []byte(string(runes))
}

func main() {
	bs := []byte("hello, world!")

	bs = reverse(bs)

	fmt.Println(string(bs))
}

package main

import (
	"fmt"
	"unicode"
)

func removeAdjacentSpaces(bs []byte) []byte {
	var out []byte
	lastSpace := 0
	for i, b := range bs {
		if unicode.IsSpace(rune(b)) {
			if lastSpace+1 == i {
				lastSpace++
				continue
			}

			lastSpace = i
			b = ' '
		}
		out = append(out, b)
	}
	return out
}

func main() {
	toChange := []byte("hello, I like \t\tyou. Nice to \n\n\nsee you!\tn")

	toChange = removeAdjacentSpaces(toChange)

	fmt.Println(string(toChange))
}

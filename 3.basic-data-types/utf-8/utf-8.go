package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r1 := '世'
	r2 := "\xe4\xb8\x96"
	r3 := '\u4e16'
	r4 := '\U00004e16'

	fmt.Println(r1 == 3)
	fmt.Println(r3 == r4)

	fmt.Println(r2)

	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(Runes(s))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%[2]d\t%[2]x\n", i, r)
	}

	fmt.Println("\x7e")

}

// HasPrefix determines whether a string s has a certain prefix.
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix determines whether a string s has a certain suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains determines whether a string s includes a certain substring s.
func Contains(s, substr string) bool {
	for i := range s {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// Runes returns a slice of the runes inside string s.
func Runes(s string) []rune {
	rns := make([]rune, 0)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		rns = append(rns, r)
		i += size
	}
	return rns
}

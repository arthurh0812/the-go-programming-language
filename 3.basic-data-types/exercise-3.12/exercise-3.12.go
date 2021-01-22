package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	if isAnagram("earlysaffenpopo", "opopneffasylrae") {
		fmt.Println("!anagram!")
	}
}

// isAnagram tests whether two strings are an anagram.
func isAnagram(str1 string, str2 string) bool {
	space := []rune(str2)
	if str1 == str2 || len([]rune(str1)) != len(space) {
		return false
	}

	var anagram bytes.Buffer

	for _, r := range str1 {
		n := strings.IndexRune(string(space), rune(r))
		if n != -1 {
			space = append(space[:n], space[n+1:]...)
			anagram.WriteRune(rune(r))
		}
	}

	return str1 == anagram.String()
}

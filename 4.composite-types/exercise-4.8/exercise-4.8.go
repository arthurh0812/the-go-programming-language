package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var categories = make(map[string]int)

func main() {
	invalid := 0

	input := bufio.NewReader(os.Stdin)

	for {
		// entire body execution for each single rune
		r, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}

		// other errors
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount2: %v\n", err)
			os.Exit(1)
		}

		// invalid character
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}

		category := ""

		if unicode.IsLetter(r) {
			category = "letter"
		} else if unicode.IsDigit(r) {
			category = "digit"
		} else if unicode.IsSpace(r) {
			category = "space"
		} else if unicode.IsSymbol(r) {
			category = "symbol"
		} else {
			category = "else"
		}

		categories[category]++
	}

	for category, count := range categories {
		if category == "else" {
			fmt.Printf("%s:\t%d\n", category, count)
			continue
		}
		fmt.Printf("%ss\t%d\n", category, count)
	}
}

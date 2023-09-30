// dup2 prints the count and text of lines that appear more than once
// in the input. It reads deom stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer file.Close()
			countLines(file, counts)
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignoring potential errors from input.Err()
}

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
			duplicates := countLines(file, counts)
			if duplicates > 0 {
				fmt.Println(filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) int {
	input := bufio.NewScanner(f)
	var sum int
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			sum++
		}
	}
	// ignoring potential errors from input.Err()
	return sum
}

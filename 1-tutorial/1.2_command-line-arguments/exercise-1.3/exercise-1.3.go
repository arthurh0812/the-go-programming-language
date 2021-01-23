package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	dur := time.Since(start).Microseconds()
	fmt.Printf("timeDelta 1: %dmics\n", dur)
	// 913 microseconds or 0 microseconds

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))

	dur = time.Since(start).Microseconds()
	fmt.Printf("timeDelta 2: %dmics\n", dur)
	// 0 microseconds
}

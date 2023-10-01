package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond * 75) // asynchronous

	n := 40
	fib := fibonacci(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fib)
}

// careful: infinite loop
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

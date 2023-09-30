package main

import (
	"fmt"

	popcount "github.com/arthurh0812/the-go-programming-language/2-program-structure/exercise-2.3/popcount"
)

func main() {
	n := popcount.PopCount(33784)

	fmt.Println(n)
}

package main

import (
	"fmt"

	popcount "github.com/arthurh0812/the-go-programming-language/2-program-structure/package-init/popcount"
)

func main() {
	n := popcount.PopCount(10446744073709551615)
	fmt.Println(n)
}

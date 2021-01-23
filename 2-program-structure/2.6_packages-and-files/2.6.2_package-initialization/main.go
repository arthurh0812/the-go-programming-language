package main

import (
	"fmt"

	"github.com/Arthur1208/private-go/the-go-programming-language/2.program-structure/2.6_packages-and-files/2.6.2_package-initialization/popcount"
)

func main() {
	n := popcount.PopCount(10446744073709551615)
	fmt.Println(n)
}

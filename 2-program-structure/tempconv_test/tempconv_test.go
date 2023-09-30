package main

import (
	"fmt"

	tempconv "github.com/arthurh0812/the-go-programming-language/2-program-structure/tempconv"
)

func main() {
	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) compile error: type mismatch
	fmt.Println(c == tempconv.Celsius(f))

	c = tempconv.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}

package main

import (
	"fmt"

	tempconv "github.com/arthurh0812/the-go-programming-language/2-program-structure/tempconv"
)

func main() {
	var c tempconv.Celsius = tempconv.Celsius(5)

	var f tempconv.Fahrenheit = tempconv.CToF(c)

	var k tempconv.Kelvin = tempconv.FToK(f)

	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(k)
}

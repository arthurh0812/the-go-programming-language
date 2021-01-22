package main

import (
	"fmt"
	"github.com/Arthur1208/private-go/the-go-programming-language/2.program-structure/exercise-2.1/tempconv"
)

func main() {
	var c tempconv.Celsius = tempconv.Celsius(5)

	var f tempconv.Fahrenheit = tempconv.CToF(c)

	var k tempconv.Kelvin = tempconv.FToK(f)

	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(k)
}

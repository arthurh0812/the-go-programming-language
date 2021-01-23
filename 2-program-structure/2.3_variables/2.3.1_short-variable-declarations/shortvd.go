package main

import (
	"fmt"
)

func main() {
	er := "ok"
	if true {
		er := "ok"
		fmt.Println(er)
		// er := "not ok"				// no new variable
		er, n := "really ok", 4 // new variable in same declaration
		fmt.Println(er, n)
	}

	fmt.Println(er)
	// er := "not ok"				// no new variable
	er, n := "really ok", 4 // new variable in same declaration
	fmt.Println(er, n)
}

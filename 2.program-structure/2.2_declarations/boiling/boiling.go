// boiling prints the boiling point of water.
// package declaration
package main

// import declaration
import (
	"fmt"
)

// package-level declaration
const boilingF = 212.0

// package-level declaration
func main() {
	// local declaration
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

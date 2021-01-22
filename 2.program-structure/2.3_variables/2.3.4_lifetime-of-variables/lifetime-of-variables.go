// lifetime-of-variables show the lifetime of variables.
package main

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

// variable path to x is not "lost" => "x escapes from func f" (outlives f)

func g() {
	y := new(int)
	*y = 1
}

// variable path to y is "lost" => memory is ready to be reclaimed

func main() {
	f()

	g()
}

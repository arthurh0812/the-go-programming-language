package main

import (
	"fmt"
)

// Point is a data structure to represent a 2D point on a coordinate system.
type Point struct {
	X, Y int
}

// Circle is a data structure to represent a Circle, with a center and a radius.
type Circle struct {
	Center Point
	Radius int
}

// Wheel is a data structure to represent a Wheel, containing of a circle and a number of spokes.
type Wheel struct {
	Circle
	Spokes int
}

var w Wheel = Wheel{Circle: Circle{Center: Point{4, 6}, Radius: 3}, Spokes: 20}

func main() {
	// short hand form
	w = Wheel{Circle{Point{10, 20}, 5}, 30}

	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Center: Point{X: 9, Y: 13},
			Radius: 3,
		},
		Spokes: 10,
	}

	fmt.Printf("%#v\n", w)

	w.Center.X = 11

	fmt.Printf("%#v\n", w)
}

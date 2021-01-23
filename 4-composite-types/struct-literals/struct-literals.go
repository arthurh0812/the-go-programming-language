package main

import (
	"fmt"
	"time"
)

// Point represents the information of a single 2D point in a coordinate system.
type Point struct {
	X, Y int
}

// Employee is the data structure which holds information about an employee.
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	p := Point{1, 2}

	fmt.Println(p)

	pe := &Employee{ID: 2357470042, Name: "Ryan Smith", Address: "London", DoB: time.Date(1995, time.January, 5, 18, 30, 0, 0, &time.Location{}), Position: "Manager", Salary: 5000, ManagerID: 44740008}

	AwardAnnualRaise(pe)
	fmt.Println(pe)

	b := Bonus(pe, 30)
	fmt.Println(b)
}

// Scale scales the coordinates of a 2D point by the provided factor.
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Bonus uses efficient change of an Employee struct using its pointer reference.
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// AwardAnnualRaise is
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

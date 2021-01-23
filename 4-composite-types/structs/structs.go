package main

import (
	"fmt"
	"log"
	"time"
)

// Employee is a struct for an employee record.
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

// Employee1 struct
type Employee1 struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	fmt.Println(dilbert.Salary, dilbert.Position)

	dilbert.Salary -= 5000

	fmt.Println(dilbert.Salary, dilbert.Position)

	position := &dilbert.Position

	*position = "Senior " + *position

	fmt.Println(dilbert.Salary, dilbert.Position)

	// with dot notation
	var employeeOfTheMonth *Employee = &dilbert

	employeeOfTheMonth.Position += " (proactive team player)"
	// equivalent to
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	EmployeeByID(dilbert.ID).Salary = 0 // fired

	anotherEmployee := Employee1(dilbert)

	log.Println(anotherEmployee)
}

// EmployeeByID returns the coorresponding employee instance whose ID matches the provided one
func EmployeeByID(id int) *Employee {
	/* access and query DB, evaluate and manipulate employee result... */
	return &Employee{}
}

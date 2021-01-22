package main

func main() {
	var x int
	var p *bool
	var person struct{ name string }
	var count []int

	scale := 3

	x = 1                       // direct assingments
	*p = true                   // indirect assignments
	person.name = "bob"         // struct field
	count[x] = count[x] * scale // array or slice or map element
}

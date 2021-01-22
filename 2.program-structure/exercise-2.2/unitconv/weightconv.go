/*Package unitconv performs general-purpose unit computations*/
package unitconv

import (
	"fmt"
)

// Gram is the data type for storing Gram weight values
type Gram float64

func (g Gram) String() string {
	return fmt.Sprintf("%gg", g)
}

// Weight method to identify weight units
func (g Gram) Weight() {

}

// Kilogram is the data type for storing Kilogram weight values
type Kilogram float64

func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}

// Weight method to identify weight units
func (kg Kilogram) Weight() {

}

// Pound is the data type for storing Pound weight values
type Pound float64

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

// Weight method to identify weight units
func (p Pound) Weight() {

}

// WeightUnit interface bounds all weight units together
type WeightUnit interface {
	Weight()
}

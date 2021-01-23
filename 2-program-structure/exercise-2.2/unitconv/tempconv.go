/*Package unitconv performs general-purpose unit computations*/
package unitconv

import (
	"fmt"
)

// Celsius is the data type for storing Celsius temperature values
type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// Temp method to identify temperature units
func (c Celsius) Temp() {

}

// Fahrenheit is the data type for storing Fahrenheit temperature values
type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// Temp method to identify temperature units
func (f Fahrenheit) Temp() {

}

// Kelvin is the data type for storing Kelvin temperature values
type Kelvin float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

// Temp method to identify temperature units
func (k Kelvin) Temp() {

}

const (
	// AbsoluteZeroC is the coldest temperature specified in Celsius
	AbsoluteZeroC Celsius = -273.15
	// AbsoluteZeroF is the coldest temperature specified in Fahrenheit
	AbsoluteZeroF Fahrenheit = -459.67
	// AbsoluteZeroK is the coldest temperature specified in Kelvin
	AbsoluteZeroK Kelvin = 0
	// FreezingC is the temperature specified in Celsius when ice starts to smelt
	FreezingC Celsius = 0
	// FreezingF is the temperature specified in Fahrenheit when ice starts to smelt
	FreezingF Fahrenheit = 32
	// FreezingK is the temperature specified in Kelvin when ice starts to smelt
	FreezingK Kelvin = 273.15
	// BoilingC is the temperature specified in Celsius when water starts to become a gas.
	BoilingC Celsius = 100
	// BoilingF is the temperature specified in Fahrenheit when water starts to become a gas.
	BoilingF Fahrenheit = 212
	// BoilingK is the temperature specified in Kelvin when water starts to become a gas.
	BoilingK Kelvin = 373.15
)

// TempUnit interface to bound all temperature units together
type TempUnit interface {
	Temp()
}

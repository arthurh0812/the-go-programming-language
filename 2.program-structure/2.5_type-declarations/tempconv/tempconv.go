/*Package tempconv performs Celsius and Fahrenheit temperature computations*/
package tempconv

import (
	"fmt"
)

// Celsius is the data type for storing Celsius temperature values.
type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// Fahrenheit is the data type for storing Fahrenheit temperature values.
type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

const (
	// AbsoluteZeroC is the coldest temperature specified in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the temperature specified in Celsius when ice starts to smelt
	FreezingC Celsius = 0
	// BoilingC is the temperature specified in Celsius when water starts to become a gas.
	BoilingC Celsius = 100
)

// CToF returns the Fahrenheit value of the given Celsius value.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC returns the Celsius value of the given Fahrenheit value.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

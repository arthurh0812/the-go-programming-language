/*Package unitconv performs general-purpose unit computations*/
package unitconv

import (
	"fmt"
)

// Centimeter is the data type for storing length values in Centimeters
type Centimeter float64

func (cm Centimeter) String() string {
	return fmt.Sprintf("%gcm", cm)
}

// Length method to identify length units
func (cm Centimeter) Length() {

}

// Meter is the data type for storing length values in Meters
type Meter float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

// Length method to identify length units
func (m Meter) Length() {

}

// Foot is the data type for storing length values in Feets
type Foot float64

func (ft Foot) String() string {
	return fmt.Sprintf("%gft", ft)
}

// Length method to identify length units
func (ft Foot) Length() {

}

// LengthUnit interface to bound all length units together
type LengthUnit interface {
	Length()
}

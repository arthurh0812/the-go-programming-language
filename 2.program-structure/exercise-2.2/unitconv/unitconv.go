/*Package unitconv performs general-purpose unit computations*/
package unitconv

// Unit defines the basic unit interface
type Unit interface {
	WeightUnit
	TempUnit
	LengthUnit
}

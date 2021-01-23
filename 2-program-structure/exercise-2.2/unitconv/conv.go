/*Package unitconv performs general-purpose unit computations*/
package unitconv

// CelsiusTo converts the given Celsius value (in the wanted Unit) to the corresponding temperature unit
func CelsiusTo(t TempUnit) TempUnit {
	if f, ok := t.(Fahrenheit); ok {
		return f*9/5 + 32
	}

	if k, ok := t.(Kelvin); ok {
		return k + 273.15
	}

	return nil
}

// FahrenheitTo converts the given Fahrenheit value (in the wanted Unit) to the corresponding temperature unit
func FahrenheitTo(t TempUnit) TempUnit {
	if c, ok := t.(Celsius); ok {
		return (c - 32) * 5 / 9
	}

	if k, ok := t.(Kelvin); ok {
		return ((k-32)*5/9 + 273.15)
	}

	return nil
}

// KelvinTo converts the given Kelvin value (in the wanted Unit) to the corresponding temperature unit
func KelvinTo(t TempUnit) TempUnit {
	if c, ok := t.(Celsius); ok {
		return c - 273.15
	}

	if f, ok := t.(Fahrenheit); ok {
		return f*9/5 - 459.67
	}

	return nil
}

// GramTo converts the given Gram value (in the wanted Unit) to the corresponding weight unit
func GramTo(w WeightUnit) WeightUnit {
	if kg, ok := w.(Kilogram); ok {
		return kg / 1000
	}

	if lb, ok := w.(Pound); ok {
		return lb / 453.59237
	}

	return nil
}

// KilogramTo converts the given Kilogram value (in the wanted Unit) to the corresponding weight unit
func KilogramTo(w WeightUnit) WeightUnit {
	if g, ok := w.(Gram); ok {
		return g * 1000
	}

	if lb, ok := w.(Pound); ok {
		return lb * 2.20462262
	}

	return nil
}

// PoundTo converts the given Pound value (in the wanted unit) to the corresponding weight unit
func PoundTo(w WeightUnit) WeightUnit {
	if g, ok := w.(Gram); ok {
		return g * 453.59237
	}

	if kg, ok := w.(Kilogram); ok {
		return kg / 2.20462262
	}

	return nil
}

// CentimeterTo converts the given Centimeter value (in the wanted unit) to the corresponding length unit
func CentimeterTo(l LengthUnit) LengthUnit {
	if m, ok := l.(Meter); ok {
		return m / 100
	}

	if ft, ok := l.(Foot); ok {
		return ft / 30.48
	}

	return nil
}

// MeterTo converts the given Meter value (in the wanted unit) to the corresponding length unit
func MeterTo(l LengthUnit) LengthUnit {
	if cm, ok := l.(Centimeter); ok {
		return cm * 100
	}

	if ft, ok := l.(Foot); ok {
		return ft * 3.2808399
	}

	return nil
}

// FootTo converts the given Centimeter value (in the wanted unit) to the corresponding length unit
func FootTo(l LengthUnit) LengthUnit {
	if cm, ok := l.(Centimeter); ok {
		return cm * 30.48
	}

	if m, ok := l.(Meter); ok {
		return m / 3.2808399
	}

	return nil
}

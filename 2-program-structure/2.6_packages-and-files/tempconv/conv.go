/*Package tempconv performs Celsius and Fahrenheit temperature computations*/
package tempconv

// CToF returns the Fahrenheit value of the given Celsius value
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK returns the Kelvin value of the given Celsius value
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToC returns the Celsius value of the given Fahrenheit value
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK returns the Kelvin value of the given Fahrenheit value
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(((f-32)*5/9 + 273.15))
}

// KToC returns the Celsius value of the given Kelvin value
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF returns the Fahrenheit value of the given Kelvin value
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}

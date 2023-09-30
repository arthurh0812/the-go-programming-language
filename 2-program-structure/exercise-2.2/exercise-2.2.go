package main

import (
	"fmt"
	"os"
	"strconv"

	unitconv "github.com/arthurh0812/the-go-programming-language/2-program-structure/exercise-2.2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		float, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise-2.1: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Temperature:\n%s  = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			unitconv.Celsius(float), unitconv.CelsiusTo(unitconv.Fahrenheit(float)),
			unitconv.Celsius(float), unitconv.CelsiusTo(unitconv.Kelvin(float)),
			unitconv.Fahrenheit(float), unitconv.FahrenheitTo(unitconv.Celsius(float)),
			unitconv.Fahrenheit(float), unitconv.FahrenheitTo(unitconv.Kelvin(float)),
			unitconv.Kelvin(float), unitconv.KelvinTo(unitconv.Celsius(float)),
			unitconv.Kelvin(float), unitconv.KelvinTo(unitconv.Fahrenheit(float)),
		)

		fmt.Printf("Weights:\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			unitconv.Gram(float), unitconv.GramTo(unitconv.Kilogram(float)),
			unitconv.Gram(float), unitconv.GramTo(unitconv.Pound(float)),
			unitconv.Kilogram(float), unitconv.KilogramTo(unitconv.Gram(float)),
			unitconv.Kilogram(float), unitconv.KilogramTo(unitconv.Pound(float)),
			unitconv.Pound(float), unitconv.PoundTo(unitconv.Gram(float)),
			unitconv.Pound(float), unitconv.PoundTo(unitconv.Kilogram(float)),
		)

		fmt.Printf("Lengths:\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			unitconv.Centimeter(float), unitconv.CentimeterTo(unitconv.Meter(float)),
			unitconv.Centimeter(float), unitconv.CentimeterTo(unitconv.Foot(float)),
			unitconv.Meter(float), unitconv.MeterTo(unitconv.Centimeter(float)),
			unitconv.Meter(float), unitconv.MeterTo(unitconv.Foot(float)),
			unitconv.Foot(float), unitconv.FootTo(unitconv.Centimeter(float)),
			unitconv.Foot(float), unitconv.FootTo(unitconv.Meter(float)),
		)
	}
}

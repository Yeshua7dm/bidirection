package main

import (
	"fmt"
	"math"
)

type Converter struct{}

type Feet float64
type Centimeter float64
type Minutes float64
type Seconds float64
type Milliseconds float64
type Celsius float64
type Fahrenheit float64
type Radian float64
type Degree float64
type Pounds float64
type Kilogram float64

func (cvr Converter) FeetToCentimeter(ft Feet) Centimeter {
	return Centimeter(ft * 30.48)
}
func (cvr Converter) CentimeterToFeet(cm Centimeter) Feet {
	return Feet(cm * 0.0328084)
}
func (cvr Converter) MinutesToSeconds(min Minutes) Seconds {
	return Seconds(min * 60)
}
func (cvr Converter) SecondsToMilliseconds(secs Seconds) Milliseconds {
	return Milliseconds(secs * 1000)
}
func (cvr Converter) CelsiusToFahrenheit(cel Celsius) Fahrenheit {
	return Fahrenheit((cel * 9 / 5) + 32)
}
func (cvr Converter) FahrenheitToCelsius(fah Fahrenheit) Celsius {
	return Celsius((fah - 32) * 5 / 9)
}
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	return Degree(rad * 180 / math.Pi)
}
func (cvr Converter) DegreeToRadian(deg Degree) Radian {
	return Radian(deg * math.Pi / 180)
}
func (cvr Converter) KilogramToPounds(kg Kilogram) Pounds {
	return Pounds(kg * 2.205)
}
func (cvr Converter) PoundsToKilogram(pds Pounds) Kilogram {
	return Kilogram(pds / 2.205)
}

func main() {
	cvr := Converter{}
	fmt.Printf("45 Centimeters is %f Feet\n", cvr.CentimeterToFeet(45))
	fmt.Printf("90 Feet is %f Centimeters\n", cvr.FeetToCentimeter(90))
	fmt.Printf("45 Minutes is %f Seconds\n", cvr.MinutesToSeconds(45))
	fmt.Printf("34 Seconds is %f Milliseconds\n", cvr.SecondsToMilliseconds(34))
	fmt.Printf("100 Degree Celsius is %f Fahrenheit\n", cvr.CelsiusToFahrenheit(100))
	fmt.Printf("34 Fahrenheit is %f Celsius\n", cvr.FahrenheitToCelsius(34))
	fmt.Printf("34 Degree is %f Feet\n", cvr.DegreeToRadian(34))
	fmt.Printf("34 Radian is %f Degree\n", cvr.RadianToDegree(34))
	fmt.Printf("10 Kilogram is %f Pounds\n", cvr.KilogramToPounds(10))
	fmt.Printf("89 Pounds is %f Kilogram\n", cvr.PoundsToKilogram(89))
}

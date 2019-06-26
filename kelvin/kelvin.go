package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celciusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celciusToFahrenheit(celsius)

	fmt.Println(kelvin, "K,", celsius, "C")
	fmt.Println(celsius, "C,", fahrenheit, "F")
}

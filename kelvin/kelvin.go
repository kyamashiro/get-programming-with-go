package main

import "fmt"

func kelvinToCelcius(k float64) float64 {
	k -= 273.15
	return k
}

func celciusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func main() {
	kelvin := 294.0
	celcius := kelvinToCelcius(kelvin)
	fahrenheit := celciusToFahrenheit(celcius)

	fmt.Println(kelvin, "K,", celcius, "C")
	fmt.Println(celcius, "C,", fahrenheit, "F")
}

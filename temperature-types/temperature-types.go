package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// レシーバ
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var f fahrenheit = 60

	c := kelvinToCelsius(k)
	fmt.Println(k, "K, ", c, "C")
	fmt.Println(k, "K, ", k.celsius(), "C")
	fmt.Println(f.celsius(), "C")
}

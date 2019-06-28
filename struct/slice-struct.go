package main

import "fmt"

func main() {
	type location struct {
		name      string
		lat, long float64
	}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.58, long: 137.4},
		{name: "Columbia", lat: -14.56, long: 175.47},
		{name: "Challenger", lat: -1.94, long: 354.47},
	}

	fmt.Println(locations)
}

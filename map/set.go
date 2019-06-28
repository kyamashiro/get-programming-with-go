package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -28.0, 10, 15,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	unique := make([]float64, 0, len(set))
	// mapは順不同
	for t := range set {
		unique = append(unique, t)
	}

	// sort前
	for _, u := range unique {
		fmt.Println(u)
	}

	fmt.Println("=== sort ===")
	// sort後
	sort.Float64s(unique)
	for _, u := range unique {
		fmt.Println(u)
	}
}

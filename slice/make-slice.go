package main

import "fmt"

func main() {
	// makeでスライスの事前割当を行う
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs, cap(dwarfs))
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris", "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs, cap(dwarfs))
}

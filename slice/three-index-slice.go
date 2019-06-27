package main

import "fmt"

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter"}
	terrestrial := planets[0:4:4]
	fmt.Println(terrestrial)
	// 新たな配列が割り当てられ、planets配列は変更されない
	worlds := append(terrestrial, "ケレス")
	fmt.Println(planets)
	fmt.Println(worlds)
}

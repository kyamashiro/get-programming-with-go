package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter"}
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	sort.Strings(planets)
	fmt.Println(planets)
	// len = length, cap = capacity
	fmt.Println(len(planets), cap(planets))
	// appendするとcapacityが2倍になる
	planets = append(planets, "add")
	fmt.Println(len(planets), cap(planets))
}

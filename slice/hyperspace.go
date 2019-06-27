package main

import (
	"fmt"
	"strings"
)

func main() {
	planets := []string{" Venus ", " Earth", "Mars "}
	hyperspace(planets)
	fmt.Println(planets)
}

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

package main

import "fmt"

func main() {
	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星",
		"冥王星",
	}

	// 配列はコピーされる
	planets2 := planets

	planets[0] = "overwrite"

	fmt.Println(planets)
	fmt.Println(planets2)

	// 値渡しなので変更されない
	terraform(planets)
	fmt.Println(planets)
}

func terraform(planets [9]string) {
	planets[1] = "overwrite"
}

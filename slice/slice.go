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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	// スライスをコピー
	iceGiantsMarkii := iceGiants

	// スライスを書き換えると基底配列(planets)も書き換わる。つまり参照渡し
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)

	fmt.Println(iceGiants, iceGiantsMarkii)
}

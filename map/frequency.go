package main

import "fmt"

func main() {
	// スライス
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -28.0,
	}

	frequency := make(map[float64]int)

	// mapのキーに気温,要素に出現回数を入れる
	// key => value みたいな
	// mapではキーの順序に保証がないので注意する
	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2fの出現は%d回です\n", t, num)
	}
}

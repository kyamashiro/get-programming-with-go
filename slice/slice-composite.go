package main

import "fmt"

func main() {
	// 複合リテラル構文による初期化
	dwarf := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// [:]でスライスできる
	dwarfSlice := dwarfArray[:]

	// arrayは固定長配列
	fmt.Printf("%T\n", dwarfArray)
	fmt.Printf("%T\n", dwarfSlice)
	fmt.Printf("%T\n", dwarf)
}

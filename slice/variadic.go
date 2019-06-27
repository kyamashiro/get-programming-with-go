package main

import "fmt"

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	// スライスで可変個引数に渡す
	planets := []string{"Venus", "Mars", "Jupiter"}
	// 引数に渡すときは省略記号...を使う
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}

// 可変個引数
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}

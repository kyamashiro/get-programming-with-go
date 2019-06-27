package main

import "fmt"

// スライスの容量の変化を表示する
func main() {
	val := make([]int, 1)
	capacity := cap(val)

	for i := 1; i < 30; i++ {
		val = append(val, i)
		if capacity != cap(val) {
			fmt.Println(cap(val))
			capacity = cap(val)
		}
	}
}

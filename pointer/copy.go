package main

import "fmt"

func main() {
	memory := 42
	memory2 := &memory
	// アドレスを共有しているので値が書き換わる
	memory = 0
	fmt.Println(memory, *memory2)

	val := 50
	val2 := &val
	val3 := *val2
	val = 0
	*val2 = 100
	// val=0xc000054088 val2=0xc000054088 val3=0xc0000540a0
	fmt.Println(&val, val2, &val3)
	// 100, 100, 50
	fmt.Println(val, val2, val3)
}

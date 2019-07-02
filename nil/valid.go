package main

import "fmt"

// 構造体の中で値の
type number struct {
	value int
	valid bool
}

// constructor
func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "未設定"
	}

	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(42)
	e := number{}

	fmt.Println(n)
	fmt.Println(e)
}

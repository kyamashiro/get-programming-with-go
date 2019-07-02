package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	// レシーバの値がnilでもメソッドが呼び出されてしまう
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()
}

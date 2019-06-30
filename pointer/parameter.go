package main

import "fmt"

func main() {
	p := person{
		name: "martin",
		age:  10,
	}
	birthday(&p)
	fmt.Println(p)

	p.birthday()
	fmt.Println(p)
}

type person struct {
	name, superpower string
	age              int
}

// メモリアドレスをデリファレンスして値を直接書き換える
// ポインタで渡すことで構造体全体がコピーされずに済む
func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

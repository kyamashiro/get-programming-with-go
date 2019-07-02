package main

import "fmt"

func main() {
	p := person{
		name: "martin",
		age:  10,
	}

	// ポインタを使用していないが自動的に変数のアドレス(&)だと判定する
	p.birthday()
	fmt.Println(p)

	birthday(&p)
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

// ポインタレシーバ
// 呼び出し元のpersonのフィールドを直接書き換えることができる
// ポインタレシーバでなければageがインクリメントされない
func (p person) birthday() {
	p.age++
}

package main

import "fmt"

func main() {
	answer := 42
	// & アドレス演算子
	// 16進数のメモリアドレス(値が保存されているRAMにの場所)が表示される
	// 0xc000054080は変数answerに割り当てられたアドレス
	fmt.Println(&answer)

	address := &answer
	// デリファレンス(逆参照・間接参照)
	// メモリアドレスで参照される値を表示する
	fmt.Println(*address)

	// 元の値42が表示される
	fmt.Println(*&answer)

	// *int型のポインタ
	fmt.Printf("addressの型は%T\n", address)

	canada := "Canada"
	var home *string
	fmt.Printf("homeは%Tです\n", home)
	home = &canada
	fmt.Printf("%+v\n", *home)
}

package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	// 存在しない要素へアクセスすると型の初期値(intなら0)が返る
	fmt.Println(temperature["nil"])

	// カンマ, ok構文
	if moon, ok := temperature["Moon"]; ok {
		fmt.Println(moon)
	} else {
		fmt.Println("none")
	}
}

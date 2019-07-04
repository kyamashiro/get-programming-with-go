package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go source(c0)
	go filter(c0, c1)
	prints(c1)
}

// ストリームの始点
func source(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		// 下流にデータを送信
		downstream <- v
	}
	// センチネル = 終了を意味する 番兵
	close(downstream)
}

// フィルター
func filter(upstream, downstream chan string) {
	// sourceから受け取ったデータをフィルタリングする
	for item := range upstream {
		// "bad"の文字列を含んでいるとき下流にデータを送る
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

// 出力
func prints(upstream chan string) {
	// 上流からきたデータを出力する
	for v := range upstream {
		fmt.Println(v)
	}
}

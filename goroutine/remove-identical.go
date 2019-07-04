package main

import "fmt"

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go upStream(c0)
	go middleStream(c0, c1)
	downStream(c1)
}

func upStream(source chan string) {
	for _, v := range []string{"a", "a", "b", "c", "d", "c"} {
		source <- v
	}
	close(source)
}

func middleStream(source, downStream chan string) {
	m := make(map[string]bool)
	// 重複を排除してチャネルを送信する
	for word := range source {
		if !m[word] {
			m[word] = true
			downStream <- word
		}
	}
	close(downStream)
}

func downStream(filtering chan string) {
	for word := range filtering {
		fmt.Println(word)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go upStream(c0)
	go split(c0, c1)
	downStream(c1)
}

func upStream(sender chan string) {
	sentence := []string{"Hello world", "Programming Go", "I send Channel"}

	for _, section := range sentence {
		sender <- section
	}

	close(sender)
}

func split(receiver, sender chan string) {
	for sentence := range receiver {
		for _, word := range strings.Fields(sentence) {
			sender <- word
		}
	}

	close(sender)
}

func downStream(receiver chan string) {
	for word := range receiver {
		fmt.Println(word)
	}
}

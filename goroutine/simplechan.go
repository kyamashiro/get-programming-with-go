package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// int型のチャネルを作成する
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGophers(i, c)
	}

	for i := 0; i < 5; i++ {
		id := <-c
		fmt.Println("gopher", id, "is awake.")
	}

}

func sleepyGophers(num int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...", num, "snore...")
	// 値を送信する
	c <- num
}

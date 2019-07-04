package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// ゴルーチンはどんな順序で実行されるかわからない
		go sleepyGopher(i)
	}

	time.Sleep(4 * time.Second)
}

func sleepyGopher(num int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", num)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(2 * time.Second)
	c := make(chan int)

	for i := 0; i < 5; i++ {
		// select文
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, "はスリープを終えました")
		case <-timeout:
			fmt.Println("忍耐の限界に達した")
			return
		}
	}
}

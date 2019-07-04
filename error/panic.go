package main

import "fmt"

func main() {
	defer func() {
		// Panicから復帰する
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("I forgot towel.")
}

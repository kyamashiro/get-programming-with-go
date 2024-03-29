package main

import (
	"fmt"
	"os"
)

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// errorを返すことで呼び出し側にエラーをどう処理するか決めさせる
func proverbs(name string) error {
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	f.Close()
	return err
}

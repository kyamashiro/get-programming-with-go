package main

import (
	"fmt"
	"os"
)

func main() {
	err := proverbs2("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverbs2(name string) error {
	f, err := os.Create(name)
	// defer 関数からリターンする前に必ず実行される
	defer f.Close()

	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}

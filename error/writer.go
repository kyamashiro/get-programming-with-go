package main

import (
	"fmt"
	"io"
	"os"
)

// エラー値をストアする
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		// エラーが発生していれば書き込みをストップ
		return
	}
	// エラーがあればストアする
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs3(name string) error {
	f, err := os.Create(name)
	defer f.Close()
	if err != nil {
		return err
	}

	sw := safeWriter{w: f}

	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")

	return sw.err
}

func main() {
	err := proverbs3("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

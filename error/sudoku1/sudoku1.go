package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

/*
エラーメッセージはユーザインタフェースの一部
エラーメッセージを代入する変数の名前は、Errではじめる
*/
var (
	// 境界外
	ErrBounds = errors.New("out of bounds")
	// 不正な数
	ErrDigit = errors.New("invalid digit")
)

func (g *Grid) Set(row, column int, digit int8) error {
	// 例外的な値をエラー処理する
	if !inBounds(row, column) {
		return ErrBounds
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if 0 < row || rows <= row {
		return false
	}

	if 0 < column || columns <= column {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	// error.Newはポインタを使って実装されているので、このswitch文はエラーメッセージのテキストではなくメモリアドレスを比較している
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Error")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}

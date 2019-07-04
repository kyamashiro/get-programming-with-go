package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

// エラーインターフェース
type SudokuError []error

func (se SudokuError) Error() string {
	var s []string

	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

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
	var errs SudokuError

	// 例外的な値をエラー処理する
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit

	// 空のerrsでもエラーと判定されるのでnilを返す
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || rows <= row {
		return false
	}

	if column < 0 || columns <= column {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	if 0 < digit && digit < 10 {
		return true
	}
	return false
}

func main() {
	var g Grid
	err := g.Set(80, 30, 11)
	// error.Newはポインタを使って実装されているので、このswitch文はエラーメッセージのテキストではなくメモリアドレスを比較している
	if err != nil {
		// 型アサーション
		// errorインターフェース型のerrの値をSudokuError型に変換する
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occured:\n", len(errs))

			for _, e := range errs {
				fmt.Printf("- %+v\n", e)
			}
		}
	}

	os.Exit(1)
}

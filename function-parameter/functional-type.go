package main

import (
	"fmt"
	"strconv"
)

// getRow関数型 返り値は文字列2つ
type getRowFn func(row int) (string, string)

// 第2パラメータはgetRow関数
func drawTable(rows int, getRow getRowFn) (string, string) {
	return getRow(rows)
}

func main() {
	var getRow getRowFn = func(row int) (s string, s2 string) {
		return strconv.Itoa(row) + "a", strconv.Itoa(row) + "b"
	}

	s1, s2 := drawTable(5, getRow)

	fmt.Println(s1, s2)
}

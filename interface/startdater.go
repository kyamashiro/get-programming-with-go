package main

import (
	"fmt"
	"time"
)

// 架空の宇宙暦の日付を計算する
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

// time.Time型がstardaterインターフェースを満たすのでtime.Dateを使うことができる
type stardater interface {
	YearDay() int
	Hour() int
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (h sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f\n", stardate(s))
}

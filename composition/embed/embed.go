package main

import "fmt"

// 構造体の埋込みをするとreport型からフィールドのメソッドに直接アクセスできる
type report struct {
	sol
	temperature // temperature型の埋込み
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64
type sol int

func main() {

	report := report{
		sol:         64,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417},
	}

	fmt.Printf("%+v\n", report.average())
	// daysメソッドの名前衝突
	fmt.Printf("%+v\n", report.days(10))
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

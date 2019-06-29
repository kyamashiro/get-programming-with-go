package main

import "fmt"

/*
階層構造の設計は難しい。また変更が広い影響に及ぶ。
小さなオブジェクトのかたまりによってより大きなオブジェクトを作成するコンポジションのほうが単純で柔軟性が高い
*/

type noCompositReport struct {
	sol       int
	high, low float64 // high,lowが何を指しているのかわからない
	lat, long float64
}

/*
report構造体は小さな構造体の集合体
より小さな型から構築しそれぞれに型にメソッドを結びつけコードを組織化する
*/
type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius // high,lowが温度を指していることがわかる
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("平均気温:%+v\n", report.temperature.average())
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

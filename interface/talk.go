package main

import (
	"fmt"
	"strings"
)

type martin struct{}

func (m martin) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

/*
interfaceによって抽象概念を表すことができる
interfaceは型が格納する値ではなく型が何を行えるかに焦点を当てる
*/
var t interface {
	talk() string
}

// interface型にはerの名前をつける
type talker interface {
	talk() string
}

// talkメソッドを持つ型はshout関数を使うことができる
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// 埋め込みでinterfaceを満たす
type starship struct {
	laser
}

func main() {
	// martinもlaserもtalkメソッドを持っているので、tに代入することができる
	t = martin{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martin{})
	shout(laser(2))

	// コンポジションとインターフェースを一緒に使うと非常に強力な設計ツールになる
	s := starship{laser(4)}
	fmt.Println(s.talk())
	shout(s)
}

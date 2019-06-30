package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

// stats構造体を書き換えるのでポインタで渡す
func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

// statsにポインタがついていないが、&player.statsで構造体内部へのポインタが得られる
// 内部ポインタとは構造体の中にあるフィールドを指し示すポインタ
type character struct {
	name  string
	stats stats
}

func main() {
	player := character{name: "Matthias"}
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)
}

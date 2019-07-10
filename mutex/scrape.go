package main

import (
	"fmt"
	"sync"
)

type visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *visited) visitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	var mu = sync.Mutex{}
	url := "https://www.google.co.jp"
	m := map[string]int{url: 0}
	v := visited{mu, m}
	count := v.visitLink(url)
	fmt.Println("count:", count)
}

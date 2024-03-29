package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

/*
*  パラメータに関数を渡す
*  func() kelvin
*  関数    返り値
 */
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor)
}

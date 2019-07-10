package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}

	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1

	for {
		select {
		// コマンドを待つ
		case c := <-r.commandc:
			switch c {
			//右に回転
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			// 左に回転
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case start:
				speed = 1
			case stop:
				speed = 0
			}
			log.Printf("new direction %v speed %d", direction, speed)
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
}

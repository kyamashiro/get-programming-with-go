package main

import (
	"fmt"
)

type stringer interface {
	String() string
}

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func addPrefix(s stringer) {
	fmt.Println("location :" + s.String())
}

func main() {
	location := location{-4.5895, 137.4417}
	addPrefix(location)
}

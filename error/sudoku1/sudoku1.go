package sudoku1

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of bounds")
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if 0 < row || rows <= row {
		return false
	}

	if 0 < column || columns <= column {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occured %+v\n", err)
		os.Exit(1)
	}
}

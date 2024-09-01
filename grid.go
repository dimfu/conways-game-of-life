package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	rows     int32
	cols     int32
	cellSize int32
	cells    [][]int32
}

func NewGrid(rows, cols, cellSize int32) *Grid {
	grid := &Grid{
		rows:     rows,
		cols:     cols,
		cellSize: cellSize,
		cells:    make([][]int32, rows),
	}

	for i := range grid.cells {
		grid.cells[i] = make([]int32, cols)
	}

	grid.FillRand()

	return grid
}

func (g Grid) isWithinBound(row, col int32) bool {
	if row >= 0 && row < g.rows && col >= 0 && col < g.cols {
		return true
	}
	return false
}

func (g *Grid) SetVal(row, col, val int32) {
	if g.isWithinBound(row, col) {
		g.cells[row][col] = val
	}
}

func (g Grid) GetVal(row, col int32) int32 {
	if g.isWithinBound(row, col) {
		return g.cells[row][col]
	}
	return 0
}

func (g *Grid) FillRand() {
	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)

	for y := 0; y < int(g.rows); y++ {
		for x := 0; x < int(g.cols); x++ {
			if rng.Float64() < 0.1 {
				g.cells[y][x] = 1
			} else {
				g.cells[y][x] = 0
			}
		}
	}
}

func (g Grid) Draw() {
	alive := rl.Color{R: 212, G: 190, B: 152, A: 255}
	dead := rl.Color{R: 60, G: 56, B: 54, A: 255}
	for row := 0; row < int(g.rows); row++ {
		for col := 0; col < int(g.cols); col++ {
			cellColor := dead
			if g.cells[row][col] == 1 {
				cellColor = alive
			}
			rl.DrawRectangle(int32(col)*g.cellSize, int32(row)*g.cellSize, g.cellSize-1, g.cellSize-1, cellColor)
		}
	}
}

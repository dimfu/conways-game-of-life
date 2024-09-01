package main

type Sim struct {
	swap           *Grid
	Grid           *Grid
	neighborCounts [][]int
}

func NewSim(rows, cols, cellSize int32) *Sim {
	sim := &Sim{
		Grid:           NewGrid(rows, cols, cellSize),
		swap:           NewGrid(rows, cols, cellSize),
		neighborCounts: make([][]int, rows),
	}

	for i := range sim.neighborCounts {
		sim.neighborCounts[i] = make([]int, cols)
	}

	return sim
}

func (s *Sim) Draw() {
	s.Grid.Draw()
}

func (s *Sim) UpdateNeighborCounts() {
	for row := int32(0); row < s.Grid.rows; row++ {
		for col := int32(0); col < s.Grid.cols; col++ {
			if s.Grid.GetVal(row, col) == 1 {
				s.incrementNeighbors(row, col)
			}
		}
	}
}

func (s *Sim) incrementNeighbors(row, col int32) {
	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nrow := (row + dx + s.Grid.rows) % s.Grid.rows
			ncol := (col + dy + s.Grid.cols) % s.Grid.cols
			s.neighborCounts[nrow][ncol]++
		}
	}
}

func (s *Sim) Update() {
	s.UpdateNeighborCounts()
	for row := 0; row < int(s.Grid.rows); row++ {
		for col := 0; col < int(s.Grid.cols); col++ {
			live := s.neighborCounts[row][col]
			cval := s.Grid.GetVal(int32(row), int32(col))

			if cval == 1 {
				if live > 3 || live < 2 {
					s.swap.SetVal(int32(row), int32(col), 0) // go ded
				} else {
					s.swap.SetVal(int32(row), int32(col), 1)
				}
			} else {
				if live == 3 {
					s.swap.SetVal(int32(row), int32(col), 1)
				} else {
					s.swap.SetVal(int32(row), int32(col), 0)
				}
			}
		}
	}

	s.Grid, s.swap = s.swap, s.Grid
	for i := range s.neighborCounts {
		for j := range s.neighborCounts[i] {
			s.neighborCounts[i][j] = 0
		}
	}
}

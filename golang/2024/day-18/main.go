package solution

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

type Coordinate struct {
	Corrupted    bool
	ShortestPath float64
	Y, X         int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var (
	OutOfMap    = fmt.Errorf("OutOfMap")
	Obstruction = fmt.Errorf("Obstruction")
)

var directions map[Direction][2]int = map[Direction][2]int{
	Up:    {-1, 0},
	Down:  {1, 0},
	Left:  {0, -1},
	Right: {0, 1},
}

type RAM struct {
	grid [][]*Coordinate
}

func createRamGrid(n int) *RAM {
	grid := make([][]*Coordinate, n)

	for i := range grid {
		grid[i] = make([]*Coordinate, n)

		for j := range grid[i] {
			grid[i][j] = &Coordinate{
				ShortestPath: math.Inf(0),
				Y:            i,
				X:            j,
			}
		}
	}

	return &RAM{
		grid: grid,
	}
}

func (r *RAM) placeCorruptedCoordinates(coordinates []string) {
	for _, coordinate := range coordinates {
		parts := strings.Split(coordinate, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		r.grid[y][x].Corrupted = true
	}
}

func (r *RAM) explore(coordinate *Coordinate) {
	path := coordinate.ShortestPath + 1

	y, x := coordinate.Y+directions[Up][0], coordinate.X+directions[Up][1]

	if (y >= 0 && y < len(r.grid)) && (x >= 0 && x < len(r.grid[0])) {
		if !r.grid[y][x].Corrupted && r.grid[y][x].ShortestPath > path {
			r.grid[y][x].ShortestPath = path

			r.explore(r.grid[y][x])
		}
	}

	y, x = coordinate.Y+directions[Down][0], coordinate.X+directions[Down][1]

	if (y >= 0 && y < len(r.grid)) && (x >= 0 && x < len(r.grid[0])) {
		if !r.grid[y][x].Corrupted && r.grid[y][x].ShortestPath > path {
			r.grid[y][x].ShortestPath = path

			r.explore(r.grid[y][x])
		}
	}

	y, x = coordinate.Y+directions[Left][0], coordinate.X+directions[Left][1]

	if (y >= 0 && y < len(r.grid)) && (x >= 0 && x < len(r.grid[0])) {
		if !r.grid[y][x].Corrupted && r.grid[y][x].ShortestPath > path {
			r.grid[y][x].ShortestPath = path

			r.explore(r.grid[y][x])
		}
	}

	y, x = coordinate.Y+directions[Right][0], coordinate.X+directions[Right][1]

	if (y >= 0 && y < len(r.grid)) && (x >= 0 && x < len(r.grid[0])) {
		if !r.grid[y][x].Corrupted && r.grid[y][x].ShortestPath > path {
			r.grid[y][x].ShortestPath = path

			r.explore(r.grid[y][x])
		}
	}
}

func (s Solution) ProcessInput(input []string) any {
	return input[:len(input)-1]
}

func (r *RAM) reset() {
	for i, row := range r.grid {
		for j, col := range row {
			if i == 0 && j == 0 {
				continue
			}
			if !col.Corrupted {
				r.grid[i][j].ShortestPath = math.Inf(0)
			}
		}
	}
}

func (s Solution) PartOne(input any) any {
	ram := createRamGrid(71)

	ram.placeCorruptedCoordinates(input.([]string)[:1024])

	ram.grid[0][0].ShortestPath = 0

	ram.explore(ram.grid[0][0])

	return ram.grid[70][70].ShortestPath
}

func (s Solution) PartTwo(input any) any {
	ram := createRamGrid(71)

	ram.placeCorruptedCoordinates(input.([]string)[:1024])
	ram.grid[0][0].ShortestPath = 0

	tests := input.([]string)[1024:]

	for _, test := range tests {
		parts := strings.Split(test, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		ram.grid[y][x].Corrupted = true

		ram.reset()

		ram.explore(ram.grid[0][0])

		if ram.grid[70][70].ShortestPath == math.Inf(0) {
			return fmt.Sprintf("%d,%d", x, y)
		}
	}

	return ram.grid[70][70].ShortestPath
}

func GetSolution() Solution {
	return Solution{}
}

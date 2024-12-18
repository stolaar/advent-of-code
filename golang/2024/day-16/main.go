package main

import (
	"math"
)

type Tile struct {
	Start, End, Visited bool
	Val                 byte
	Y, X                int
	Path                float64
}

type Maze struct {
	grid  [][]*Tile
	Start *Tile
	End   *Tile
}

func linesToMaze(lines []string) *Maze {
	maze, grid := &Maze{}, make([][]*Tile, len(lines)-1)

	for i := range grid {
		grid[i] = make([]*Tile, len(lines[0]))

		for j := range grid[i] {
			tile := &Tile{
				Val:   lines[i][j],
				Y:     i,
				X:     j,
				Start: lines[i][j] == 'S',
				End:   lines[i][j] == 'E',
				Path:  math.Inf(0),
			}

			if tile.Start {
				tile.Path = 0
				maze.Start = tile
			}

			if tile.End {
				maze.End = tile
			}

			grid[i][j] = tile
		}
	}

	maze.grid = grid

	return maze
}

func ProcessInput(input []string) interface{} {
	return input
}

var movementsMap map[byte][2]int = map[byte][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
}

var rotationsCosts map[byte]map[byte]int = map[byte]map[byte]int{
	'^': {
		'>': 1000,
		'<': 1000,
		'v': 2000,
	},
	'>': {
		'^': 1000,
		'<': 2000,
		'v': 1000,
	},
	'<': {
		'^': 1000,
		'>': 2000,
		'v': 1000,
	},
	'v': {
		'^': 2000,
		'>': 1000,
		'<': 1000,
	},
}

type Element struct {
	tile      *Tile
	cost      int
	direction byte
}

func (m *Maze) explore(tile *Tile, direction byte) {
	y, x := tile.Y+movementsMap[direction][0], tile.X+movementsMap[direction][1]

	nextTile := m.grid[y][x]

	if nextTile.Val != '#' && nextTile.Path > tile.Path+1 {
		nextTile.Path = tile.Path + 1

		m.explore(nextTile, direction)
	}

	for key, val := range rotationsCosts[direction] {
		y, x := tile.Y+movementsMap[key][0], tile.X+movementsMap[key][1]

		if y < 0 || y >= len(m.grid) || x < 0 || x >= len(m.grid[0]) {
			continue
		}

		nextTile := m.grid[y][x]

		if nextTile.Val != '#' && nextTile.Path > tile.Path+float64(val)+1 {
			nextTile.Path = tile.Path + float64(val) + 1

			m.explore(nextTile, key)
		}
	}
}

func PartOne(input interface{}) interface{} {
	m := linesToMaze(input.([]string))

	m.explore(m.Start, '>')

	return m.End.Path
}

func PartTwo(input interface{}) interface{} {
	return ""
}

package main

import (
	"strconv"
)

func ProcessInput(input []string) interface{} {
	grid := make([][]*Trail, len(input)-1)

	for i := 0; i < len(input[0]); i++ {
		grid[i] = make([]*Trail, len(input[0]))

		for j := range grid[i] {
			val, _ := strconv.Atoi(string(input[i][j]))

			grid[i][j] = &Trail{
				Val: val,
				X:   j,
				Y:   i,
			}
		}
	}

	return &TopographicMap{
		grid:       grid,
		totalScore: 0,
	}
}

type Trail struct {
	Val, Y, X int
}

type TopographicMap struct {
	grid       [][]*Trail
	totalScore int
}

func (tm *TopographicMap) explore(headTrail *Trail, distinct bool) int {
	q, reached := []*Trail{headTrail}, map[*Trail]bool{}

	sum := 0

	for len(q) > 0 {
		trail := q[0]
		q = q[1:]

		if trail.Val == 9 {
			if _, ok := reached[trail]; ok && distinct {
				continue
			}
			reached[trail] = true
			sum += 1
			continue
		}

		if trail.Y > 0 && tm.grid[trail.Y-1][trail.X].Val == trail.Val+1 {
			q = append(q, tm.grid[trail.Y-1][trail.X])
		}

		if trail.Y < len(tm.grid)-1 && tm.grid[trail.Y+1][trail.X].Val == trail.Val+1 {
			q = append(q, tm.grid[trail.Y+1][trail.X])
		}

		if trail.X > 0 && tm.grid[trail.Y][trail.X-1].Val == trail.Val+1 {
			q = append(q, tm.grid[trail.Y][trail.X-1])
		}

		if trail.X < len(tm.grid[0])-1 && tm.grid[trail.Y][trail.X+1].Val == trail.Val+1 {
			q = append(q, tm.grid[trail.Y][trail.X+1])
		}
	}

	return sum
}

func PartOne(input interface{}) interface{} {
	tp, sum := input.(*TopographicMap), 0

	for _, row := range tp.grid {
		for _, col := range row {
			if col.Val == 0 {
				sum += tp.explore(col, true)
			}
		}
	}

	return sum
}

func PartTwo(input interface{}) interface{} {
	tp, sum := input.(*TopographicMap), 0

	for _, row := range tp.grid {
		for _, col := range row {
			if col.Val == 0 {
				sum += tp.explore(col, false)
			}
		}
	}

	return sum
}

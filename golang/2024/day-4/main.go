package main

import (
	"strings"
)

type Column struct {
	Val byte
	X   int
	Y   int
}

type WordSearch struct {
	grid [][]*Column
}

func ProcessInput(input []string) interface{} {
	grid := make([][]*Column, len(input)-1)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]*Column, len(input[0]))

		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = &Column{
				Val: input[i][j],
				X:   j,
				Y:   i,
			}
		}
	}

	return &WordSearch{
		grid: grid,
	}
}

func (ws *WordSearch) lookUp(column *Column) bool {
	var str strings.Builder

	if column.Y-3 < 0 {
		return false
	}

	for i := column.Y - 1; i >= max(0, column.Y-3); i-- {
		s := ws.grid[i][column.X].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookDown(column *Column) bool {
	var str strings.Builder

	if column.Y+3 >= len(ws.grid) {
		return false
	}

	for i := column.Y + 1; i < min(len(ws.grid), column.Y+4); i++ {
		s := ws.grid[i][column.X].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookRight(column *Column) bool {
	var str strings.Builder

	if column.X+3 >= len(ws.grid[0]) {
		return false
	}

	for j := column.X + 1; j < min(len(ws.grid[0]), column.X+4); j++ {
		s := ws.grid[column.Y][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookLeft(column *Column) bool {
	var str strings.Builder

	if column.X-3 < 0 {
		return false
	}

	for j := column.X - 1; j >= max(0, column.X-3); j-- {
		s := ws.grid[column.Y][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookLeftUp(column *Column) bool {
	var str strings.Builder

	if column.X-3 < 0 || column.Y-3 < 0 {
		return false
	}

	k := column.Y - 1

	for j := column.X - 1; j >= max(0, column.X-3); j-- {
		s := ws.grid[k][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		k--
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookRightUp(column *Column) bool {
	var str strings.Builder

	if column.X+3 >= len(ws.grid[0]) || column.Y-3 < 0 {
		return false
	}

	k := column.Y - 1

	for j := column.X + 1; j < min(len(ws.grid[0]), column.X+4); j++ {
		s := ws.grid[k][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		k--
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookLeftDown(column *Column) bool {
	var str strings.Builder

	if column.X-3 < 0 || column.Y+3 >= len(ws.grid) {
		return false
	}

	k := column.Y + 1

	for j := column.X - 1; j >= max(0, column.X-3); j-- {
		s := ws.grid[k][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		k++
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) lookRightDown(column *Column) bool {
	var str strings.Builder

	if column.X+3 >= len(ws.grid[0]) || column.Y+3 >= len(ws.grid) {
		return false
	}

	k := column.Y + 1

	for j := column.X + 1; j < min(len(ws.grid[0]), column.X+4); j++ {
		s := ws.grid[k][j].Val

		if s != 'M' && s != 'A' && s != 'S' {
			break
		}
		k++
		str.WriteByte(s)
	}

	return str.String() == "MAS"
}

func (ws *WordSearch) countOccurences(column *Column) int {
	count := 0

	if ws.lookUp(column) {
		count += 1
	}

	if ws.lookDown(column) {
		count += 1
	}

	if ws.lookLeft(column) {
		count += 1
	}

	if ws.lookRight(column) {
		count += 1
	}

	if ws.lookLeftUp(column) {
		count += 1
	}

	if ws.lookLeftDown(column) {
		count += 1
	}

	if ws.lookRightUp(column) {
		count += 1
	}

	if ws.lookRightDown(column) {
		count += 1
	}
	return count
}

func (ws *WordSearch) isXMas(column *Column) int {
	if column.X == 0 || column.X == len(ws.grid[0])-1 {
		return 0
	}

	if column.Y == 0 || column.Y == len(ws.grid)-1 {
		return 0
	}

	lt, rt, lb, rb := ws.grid[column.Y-1][column.X-1].Val, ws.grid[column.Y-1][column.X+1].Val, ws.grid[column.Y+1][column.X-1].Val, ws.grid[column.Y+1][column.X+1].Val

	leftOk := (lt == 'M' && rb == 'S') || (lt == 'S' && rb == 'M')
	rightOk := (rt == 'M' && lb == 'S') || (rt == 'S' && lb == 'M')

	if leftOk && rightOk {
		return 1
	}

	return 0
}

func PartOne(input interface{}) interface{} {
	ws, total := input.(*WordSearch), 0

	for _, row := range ws.grid {
		for _, col := range row {
			if col.Val == 'X' {
				total += ws.countOccurences(col)
			}
		}
	}
	return total
}

func PartTwo(input interface{}) interface{} {
	ws, total := input.(*WordSearch), 0

	for _, row := range ws.grid {
		for _, col := range row {
			if col.Val == 'A' {
				total += ws.isXMas(col)
			}
		}
	}
	return total
}

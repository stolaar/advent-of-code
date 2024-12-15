package main

import "fmt"

type Column struct {
	Val  byte
	Y, X int
	Wall bool
}

var movementsMap map[byte][2]int = map[byte][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
}

type MovementSeq struct {
	K int
	V byte
}

type Warehouse struct {
	grid      [][]*Column
	movements []byte
	position  [2]int
}

func linesToWarehouse(lines []string, doubleGrid bool) *Warehouse {
	warehouse := &Warehouse{}

	grid, gridProcessed := [][]*Column{}, false
	movements := []byte{}

	for idx, line := range lines {
		if idx == len(lines)-1 {
			break
		}

		if line == "" {
			gridProcessed = true
			continue
		}

		if !gridProcessed {
			var width int = len(line)

			if doubleGrid {
				width *= 2
			}

			col := make([]*Column, width)

			k := 0

			for j, r := range line {
				cell := &Column{
					Val:  byte(r),
					Y:    idx,
					X:    j,
					Wall: r == '#',
				}

				if doubleGrid {
					if r == '@' {
						warehouse.position = [2]int{idx, k}
						col[k] = cell
						cell.Val = '.'
						col[k+1] = cell
						k += 2
						continue
					}

					if r == '#' || r == '.' {
						col[k] = cell
						col[k+1] = cell
						k += 2
						continue
					}

					cell.Val = '['
					col[k] = cell
					cell.Val = ']'
					col[k+1] = cell
					k += 2

					continue
				}

				col[j] = cell
				if r == '@' {
					warehouse.position = [2]int{idx, j}
				}
			}
			grid = append(grid, col)
			continue
		}

		for j := range line {
			movements = append(movements, line[j])
		}
	}
	warehouse.movements = movements
	warehouse.grid = grid

	return warehouse
}

func ProcessInput(input []string) interface{} {
	return input
}

func groupMovements(movements []byte) []MovementSeq {
	groups := []MovementSeq{}

	i, j := 0, 1

	for j < len(movements) {
		streak := 1

		for k := j; k < len(movements); k++ {
			if movements[k] != movements[i] {
				j = k
				break
			}
			streak++
		}

		group := MovementSeq{streak, movements[i]}
		groups = append(groups, group)

		i = j
		j += 1
	}

	return groups
}

func (w *Warehouse) move(movement MovementSeq) {
	y, x := w.position[0], w.position[1]

	for i := 0; i < movement.K; i++ {
		y += movementsMap[movement.V][0]
		x += movementsMap[movement.V][1]

		currentCell := w.grid[y][x]

		if currentCell.Wall {
			break
		}

		if currentCell.Val == '.' {
			w.grid[w.position[0]][w.position[1]].Val = '.'
			w.position[0] = y
			w.position[1] = x
			w.grid[y][x].Val = '@'
			continue
		}

		moved := false
		for nextY, nextX := y+movementsMap[movement.V][0], x+movementsMap[movement.V][1]; !w.grid[nextY][nextX].Wall; nextY, nextX = nextY+movementsMap[movement.V][0], nextX+movementsMap[movement.V][1] {
			if w.grid[nextY][nextX].Val == '.' {
				w.grid[nextY][nextX].Val = 'O'
				moved = true
				break
			}
		}

		if moved {
			w.grid[w.position[0]][w.position[1]].Val = '.'
			w.position[0] = y
			w.position[1] = x
			w.grid[y][x].Val = '@'
			continue
		}
		break
	}
}

func printMap(grid [][]*Column) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col.Val))
		}
		println()
	}
}

func PartOne(input interface{}) interface{} {
	warehouse := linesToWarehouse(input.([]string), false)

	movements := groupMovements(warehouse.movements)

	for _, movement := range movements {
		warehouse.move(movement)
	}

	total := 0

	for i, row := range warehouse.grid {
		for j, col := range row {
			if col.Val == 'O' {
				total += (100 * i) + j
			}
		}
	}

	return total
}

func PartTwo(input interface{}) interface{} {
	return ""
}


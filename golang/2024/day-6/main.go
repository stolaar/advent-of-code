package solution

import (
	"fmt"
)

type Solution struct{}

type Column struct {
	Val      byte
	X        int
	Y        int
	Visited  bool
	Starting bool
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

func (direction Direction) next(column *Column, grid [][]*Column, testObstacle *Column) (*Column, error) {
	y, x := column.Y+directions[direction][0], column.X+directions[direction][1]

	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return nil, OutOfMap
	}

	nextColumn := grid[y][x]

	if nextColumn == nil {
		return nil, OutOfMap
	}

	if nextColumn.Val == '#' || nextColumn == testObstacle {
		return nextColumn, Obstruction
	}

	return nextColumn, nil
}

type Guard struct {
	grid                [][]*Column
	position            [2]int
	totalDistinctVisits int
	direction           Direction
	startingPosition    [2]int
}

func (guard *Guard) move(testObstacle *Column) (*Column, error) {
	next, err := guard.direction.next(guard.grid[guard.position[0]][guard.position[1]], guard.grid, testObstacle)
	if err != nil {
		if err == OutOfMap {
			return nil, err
		}

		switch guard.direction {
		case Up:
			guard.direction = Right
		case Down:
			guard.direction = Left
		case Right:
			guard.direction = Down
		case Left:
			guard.direction = Up
		}
		return next, err
	}

	return next, nil
}

func (s Solution) ProcessInput(input []string) any {
	grid, guard := make([][]*Column, len(input)-1), &Guard{
		direction: Up,
	}

	for i := 0; i < len(input[0]); i++ {
		grid[i] = make([]*Column, len(input[0]))

		for j := range grid[i] {
			isGuardPosition := input[i][j] == '^'

			grid[i][j] = &Column{
				Val:      input[i][j],
				X:        j,
				Y:        i,
				Visited:  isGuardPosition,
				Starting: isGuardPosition,
			}

			if isGuardPosition {
				guard.position = [2]int{i, j}
				guard.startingPosition = [2]int{i, j}
				guard.totalDistinctVisits = 1
			}
		}
	}

	guard.grid = grid

	return guard
}

func (s Solution) PartOne(input any) any {
	guard := input.(*Guard)

	for {
		next, err := guard.move(nil)

		if err == nil {
			guard.position = [2]int{next.Y, next.X}
			if !next.Visited {
				guard.totalDistinctVisits += 1
			}
			next.Visited = true
		}

		if err == OutOfMap {
			return guard.totalDistinctVisits
		}
	}
}

func (s Solution) PartTwo(input any) any {
	guard, testObstacles := input.(*Guard), []*Column{}

	for _, row := range guard.grid {
		for _, column := range row {
			if !column.Starting && column.Visited {
				testObstacles = append(testObstacles, column)
			}
		}
	}

	sum := 0

	for i := 0; i < len(testObstacles); i++ {
		guard.position = guard.startingPosition
		guard.direction = Up

		obstructionsCount := make(map[*Column]int)
		col := testObstacles[i]

		for {
			next, err := guard.move(col)

			if err == nil {
				guard.position = [2]int{next.Y, next.X}
				continue
			}

			if err == OutOfMap {
				break
			}

			if err == Obstruction {
				if obstructionsCount[next] > 2 {
					sum += 1
					break
				}
				obstructionsCount[next] += 1
			}
		}
	}

	return sum
}

func GetSolution() Solution {
	return Solution{}
}

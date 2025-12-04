package solution

import (
	"math"
)

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	grid, antennas := make([][]*Column, len(input)-1), map[byte][]*Column{}

	for i := 0; i < len(input)-1; i++ {
		grid[i] = make([]*Column, len(input[0]))

		for j := range grid[i] {
			val := input[i][j]

			grid[i][j] = &Column{
				Val:      val,
				X:        j,
				Y:        i,
				Antinode: false,
			}

			if val != '.' {
				antennas[val] = append(antennas[val], grid[i][j])
			}
		}
	}

	return AntennasMap{
		grid:     grid,
		antennas: antennas,
	}
}

type Column struct {
	Val      byte
	X        int
	Y        int
	Antinode bool
}

type AntennasMap struct {
	grid     [][]*Column
	antennas map[byte][]*Column
}

func (am *AntennasMap) placeAntinodes(antenna1, antenna2 *Column, atAnyPlace bool) {
	if atAnyPlace {
		antenna1.Antinode = true
		antenna2.Antinode = true
	}
	yDiff, xDiff := int(math.Abs(float64(antenna1.Y-antenna2.Y))), int(math.Abs(float64(antenna1.X-antenna2.X)))

	if antenna1.Y == antenna2.Y {
		offset := max(1, xDiff*2)
		left, right := min(antenna1.X, antenna2.X), max(antenna1.X, antenna2.X)
		leftX, rightX := left-offset, right+offset

		if leftX >= 0 {
			am.grid[antenna1.Y][leftX].Antinode = true
			if atAnyPlace {
				leftX -= offset
				for leftX >= 0 {
					am.grid[antenna1.Y][leftX].Antinode = true
					leftX -= offset
				}
			}
		}

		if rightX < len(am.grid[0]) {
			am.grid[antenna1.Y][rightX].Antinode = true

			if atAnyPlace {
				rightX += offset
				for rightX < len(am.grid[0]) {
					am.grid[antenna1.Y][rightX].Antinode = true
					rightX += offset
				}
			}
		}

		return
	}

	if antenna1.X == antenna2.X {
		offset := max(1, yDiff*2)
		top, bottom := min(antenna1.Y, antenna2.Y), max(antenna1.Y, antenna2.Y)
		topY, bottomY := top-offset, bottom+offset

		if topY >= 0 {
			am.grid[topY][antenna1.X].Antinode = true

			if atAnyPlace {
				topY -= offset
				for topY >= 0 {
					am.grid[topY][antenna1.X].Antinode = true
					topY -= offset
				}
			}
		}

		if bottomY < len(am.grid) {
			am.grid[bottomY][antenna1.X].Antinode = true

			if atAnyPlace {
				bottomY += offset
				for bottomY < len(am.grid) {
					am.grid[bottomY][antenna1.X].Antinode = true
					bottomY += offset
				}
			}
		}

		return
	}

	if antenna1.X < antenna2.X && antenna1.Y < antenna2.Y {
		left, top := antenna1.X-max(1, xDiff), antenna1.Y-max(1, yDiff)

		if left >= 0 && top >= 0 {
			am.grid[top][left].Antinode = true
			if atAnyPlace {
				left -= max(1, xDiff)
				top -= max(1, yDiff)
				for left >= 0 && top >= 0 {
					am.grid[top][left].Antinode = true
					left -= max(1, xDiff)
					top -= max(1, yDiff)
				}
			}
		}

		right, bottom := antenna2.X+max(1, xDiff), antenna2.Y+max(1, yDiff)

		if right < len(am.grid[0]) && bottom < len(am.grid) {
			am.grid[bottom][right].Antinode = true
			if atAnyPlace {
				right += max(1, xDiff)
				bottom += max(1, yDiff)
				for right < len(am.grid[0]) && bottom < len(am.grid) {
					am.grid[bottom][right].Antinode = true
					right += max(1, xDiff)
					bottom += max(1, yDiff)
				}
			}
		}
		return
	}

	if antenna1.X < antenna2.X && antenna1.Y > antenna2.Y {
		right, top := antenna2.X+max(1, xDiff), antenna2.Y-max(1, yDiff)
		if right < len(am.grid[0]) && top >= 0 {
			am.grid[top][right].Antinode = true
			if atAnyPlace {
				right += max(1, xDiff)
				top -= max(1, yDiff)
				for right < len(am.grid[0]) && top >= 0 {
					am.grid[top][right].Antinode = true
					right += max(1, xDiff)
					top -= max(1, yDiff)
				}
			}
		}

		left, bottom := antenna1.X-max(1, xDiff), antenna1.Y+min(1, yDiff)

		if left >= 0 && bottom < len(am.grid) {
			am.grid[bottom][left].Antinode = true

			if atAnyPlace {
				left -= max(1, xDiff)
				bottom += max(1, yDiff)
				for left >= 0 && bottom < len(am.grid) {
					am.grid[bottom][left].Antinode = true
					left -= max(1, xDiff)
					bottom += max(1, yDiff)
				}
			}
		}
		return
	}

	if antenna1.X > antenna2.X && antenna1.Y < antenna2.Y {
		right, top := antenna1.X+max(1, xDiff), antenna1.Y-max(1, yDiff)
		if right < len(am.grid[0]) && top >= 0 {
			am.grid[top][right].Antinode = true

			if atAnyPlace {
				right += max(1, xDiff)
				top -= max(1, yDiff)
				for right < len(am.grid[0]) && top >= 0 {
					am.grid[top][right].Antinode = true
					right += max(1, xDiff)
					top -= max(1, yDiff)
				}
			}
		}

		left, bottom := antenna2.X-(max(1, xDiff)), antenna2.Y+max(1, yDiff)

		if left >= 0 && bottom < len(am.grid) {
			am.grid[bottom][left].Antinode = true
			if atAnyPlace {
				left -= max(1, xDiff)
				bottom += max(1, yDiff)
				for left >= 0 && bottom < len(am.grid) {
					am.grid[bottom][left].Antinode = true
					left -= max(1, xDiff)
					bottom += max(1, yDiff)
				}
			}
		}
		return
	}

	if antenna1.X > antenna2.X && antenna1.Y > antenna2.Y {
		left, top := antenna2.X-max(1, xDiff), antenna2.Y-max(1, yDiff)

		if left >= 0 && top >= 0 {
			am.grid[top][left].Antinode = true
			if atAnyPlace {
				left -= max(1, xDiff)
				top -= max(1, yDiff)
				for left >= 0 && top >= 0 {
					am.grid[top][left].Antinode = true
					left -= max(1, xDiff)
					top -= max(1, yDiff)
				}
			}
		}

		right, bottom := antenna1.X+max(1, xDiff), antenna1.Y+max(1, yDiff)

		if right < len(am.grid[0]) && bottom < len(am.grid) {
			am.grid[bottom][left].Antinode = true

			if atAnyPlace {
				right += max(1, xDiff)
				bottom += max(1, yDiff)
				for right < len(am.grid[0]) && bottom < len(am.grid) {
					am.grid[top][left].Antinode = true
					right += max(1, xDiff)
					bottom += max(1, yDiff)
				}
			}
		}
		return
	}
}

func (s Solution) PartOne(input any) any {
	antennasMap := input.(AntennasMap)

	for _, antennas := range antennasMap.antennas {
		for i := 0; i < len(antennas); i++ {
			k := i + 1

			for k < len(antennas) {
				antennasMap.placeAntinodes(antennas[i], antennas[k], false)
				k++
			}
		}
	}

	sum := 0

	for _, row := range antennasMap.grid {
		for _, col := range row {
			if col.Antinode {
				sum += 1
			}
		}
	}

	return sum
}

func (s Solution) PartTwo(input any) any {
	antennasMap := input.(AntennasMap)

	for _, antennas := range antennasMap.antennas {
		for i := 0; i < len(antennas); i++ {
			k := i + 1

			for k < len(antennas) {
				antennasMap.placeAntinodes(antennas[i], antennas[k], true)
				k++
			}
		}
	}

	sum := 0

	for _, row := range antennasMap.grid {
		for _, col := range row {
			if col.Antinode {
				sum += 1
			}
		}
	}

	return sum
}


func GetSolution() Solution {
	return Solution{}
}

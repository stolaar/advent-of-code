package solution

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}

type point struct {
	x, y int
}

func (s Solution) ReProcessInput() bool {
	return false
}

func (s Solution) ProcessInput(input []string) any {
	dst := make([]point, len(input))

	for i, line := range input {
		pointsStrArr := strings.Split(line, ",")
		x, _ := strconv.Atoi(pointsStrArr[0])
		y, _ := strconv.Atoi(pointsStrArr[1])

		p := point{x, y}
		dst[i] = p
	}
	return dst
}

func (s Solution) PartOne(input any) any {
	puzzle := input.([]point)

	areas := []int{}

	for i, p := range puzzle {
		for j := i + 1; j < len(puzzle); j++ {
			horizontalLen := int(math.Abs((float64(p.x))-float64(puzzle[j].x))) + 1
			verticalLen := int(math.Abs((float64(p.y))-float64(puzzle[j].y))) + 1

			areas = append(areas, horizontalLen*verticalLen)
		}
	}
	return slices.Max(areas)
}

func (s Solution) PartTwo(input any) any {
	puzzle := input.([]point)

	xBoundaries := map[int][2]int{}

	for _, p := range puzzle {
		if _, ok := xBoundaries[p.y]; !ok {
			xBoundaries[p.y] = [2]int{p.x, p.x}
		}

		row := xBoundaries[p.y]

		if p.x < row[0] {
			row[0] = p.x
		}

		if p.x > row[1] {
			row[1] = p.x
		}

		xBoundaries[p.y] = row
	}

	maxArea := 0
	for _, p := range puzzle {
		for j := range puzzle {
			pp := puzzle[j]
			horizontalLen := int(math.Abs((float64(p.x))-float64(puzzle[j].x))) + 1
			verticalLen := int(math.Abs((float64(p.y))-float64(puzzle[j].y))) + 1

			area := horizontalLen * verticalLen

			if area < maxArea {
				continue
			}

			horizontalBoundary := [2]int{min(p.x, pp.x), max(p.x, pp.x)}

			valid := false

			if p.x > pp.x && p.y < pp.y {
				valid = topRight(xBoundaries[p.y], horizontalBoundary, xBoundaries, p) && bottomLeft(xBoundaries[pp.y], horizontalBoundary, xBoundaries, pp)
			}

			if p.x > pp.x && p.y > pp.y {
				valid = bottomRight(xBoundaries[p.y], horizontalBoundary, xBoundaries, p) && topLeft(xBoundaries[pp.y], horizontalBoundary, xBoundaries, pp)
			}

			if p.x < pp.x && p.y < pp.y {
				valid = topLeft(xBoundaries[p.y], horizontalBoundary, xBoundaries, p) && bottomRight(xBoundaries[pp.y], horizontalBoundary, xBoundaries, pp)
			}

			if p.x < pp.x && p.y > pp.y {
				valid = bottomLeft(xBoundaries[p.y], horizontalBoundary, xBoundaries, p) && topRight(xBoundaries[pp.y], horizontalBoundary, xBoundaries, pp)
			}

			if !valid {
				continue
			}

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func topRight(rowSpan [2]int, requiredSpan [2]int, xBoundaries map[int][2]int, p point) bool {
	if rowSpan[0] <= requiredSpan[0] {
		return true
	}

	for y, boundaries := range xBoundaries {
		if y >= p.y {
			continue
		}
		if boundaries[0] <= requiredSpan[0] && boundaries[1] >= requiredSpan[0] && boundaries[0] != boundaries[1] {
			return true
		}
	}

	return false
}

func topLeft(rowSpan [2]int, requiredSpan [2]int, xBoundaries map[int][2]int, p point) bool {
	if rowSpan[1] >= requiredSpan[1] {
		return true
	}

	for y, boundaries := range xBoundaries {
		if y >= p.y {
			continue
		}
		if boundaries[0] <= requiredSpan[0] && boundaries[1] >= requiredSpan[1] && boundaries[0] != boundaries[1] {
			return true
		}
	}

	return false
}

func bottomRight(rowSpan [2]int, requiredSpan [2]int, xBoundaries map[int][2]int, p point) bool {
	if rowSpan[0] <= requiredSpan[0] {
		return true
	}

	for y, boundaries := range xBoundaries {
		if y <= p.y {
			continue
		}
		if boundaries[0] <= requiredSpan[0] && boundaries[1] >= requiredSpan[0] && boundaries[0] != boundaries[1] {
			return true
		}
	}

	return false
}
func bottomLeft(rowSpan [2]int, requiredSpan [2]int, xBoundaries map[int][2]int, p point) bool {
	if rowSpan[1] >= requiredSpan[1] {
		return true
	}

	for y, boundaries := range xBoundaries {
		if y <= p.y {
			continue
		}
		if boundaries[1] >= requiredSpan[1] && boundaries[0] <= requiredSpan[1] && boundaries[0] != boundaries[1] {
			return true
		}
	}

	return false
}

func GetSolution() Solution {
	return Solution{}
}

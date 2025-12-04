package solution

import (
	"math"
	"strconv"
	"strings"
)

type Direction int

const (
	Idle Direction = iota
	Increasing
	Decreasing
)

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	reports := make([][]int, len(input)-1)

	for idx, line := range input {
		if idx == len(input)-1 {
			break
		}

		levelsAsStrings := strings.Fields(line)
		levels := make([]int, len(levelsAsStrings))

		for idx, levelAsString := range levelsAsStrings {
			level, _ := strconv.Atoi(levelAsString)
			levels[idx] = level
		}

		reports[idx] = levels
	}
	return reports
}

func getDirection(a, b int) Direction {
	if a > b {
		return Decreasing
	}

	return Increasing
}

func isValid(a, b int, direction Direction) (bool, Direction) {
	if a == b {
		return false, direction
	}

	if direction == Idle {
		direction = getDirection(a, b)
	}

	if direction == Decreasing && a < b {
		return false, direction
	}

	if direction == Increasing && a > b {
		return false, direction
	}

	diff := getDiff(a, b)

	if diff > 3 {
		return false, direction
	}
	return true, direction
}

func (s Solution) PartOne(input any) any {
	reports, result := input.([][]int), 0

	for _, levels := range reports {
		valid := true

		i, j, direction := 0, 1, Idle

		for j < len(levels) {
			valid, direction = isValid(levels[i], levels[j], direction)

			if !valid {
				break
			}
			i += 1
			j += 1
		}

		if valid {
			result += 1
		}

	}

	return result
}

func getDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func checkLevels(levels []int, initialCheck bool) bool {
	i, j, direction, valid := 0, 1, Idle, true

	for j < len(levels) {
		valid, direction = isValid(levels[i], levels[j], direction)

		if !valid {
			if !initialCheck {
				return false
			}

			l := make([]int, len(levels))
			r := make([]int, len(levels))
			r2 := make([]int, len(levels))
			copy(l, levels)
			copy(r, levels)
			copy(r2, levels)

			valid = checkLevels(append(l[0:i], l[j:]...), false) || checkLevels(append(levels[:j], levels[j+1:]...), false)

			if valid {
				return true
			}

			if !valid && i > 0 {
				valid = checkLevels(append(r[:i-1], r[i:]...), false)
			}

			if !valid && j < len(levels)-1 {
				return checkLevels(append(r[:j+1], r[j+2:]...), false)
			}

			if valid {
				return true
			}
		}

		i += 1
		j += 1
	}

	return valid
}

func (s Solution) PartTwo(input any) any {
	reports, result := input.([][]int), 0

	for _, levels := range reports {
		valid := checkLevels(levels, true)
		if valid {
			result += 1
		}
	}

	return result
}

func GetSolution() Solution {
	return Solution{}
}

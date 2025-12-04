package solution

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) ReProcessInput() bool {
	return false
}

func (s Solution) ProcessInput(input []string) any {
	return input[0 : len(input)-1]
}

func rotationsPassedZero(current, rotation int) int {
	initialPosition := current
	current += rotation

	if current > 0 && current <= 99 {
		return 0
	}

	if current == 0 {
		return 1
	}

	if current < 0 {
		diff := ((rotation * -1) - initialPosition) % 100
		if diff < 0 {
			diff *= -1
		}
		passed := max((((rotation*-1)-initialPosition)/100)+1, 1)

		current = 100 - diff

		if initialPosition == 0 {
			passed--
		}

		return max(passed, 0)
	}

	passed := max(current/100, 1)

	current = ((current) % 100)

	return max(passed, 0)
}

func rotate(current, rotation int) int {
	initialPosition := current
	current += rotation

	if current < 0 {
		diff := ((rotation * -1) - initialPosition) % 100
		if diff < 0 {
			diff *= -1
		}
		current = 100 - diff
	}

	if current > 99 {
		current = ((current) % 100)
	}
	return current
}

func (s Solution) PartOne(input any) any {
	rotations := input.([]string)

	ans, current := 0, 50

	for _, line := range rotations {
		arr := strings.FieldsFunc(line, func(r rune) bool {
			return r == 'L' || r == 'R'
		})
		distance, _ := strconv.Atoi(arr[0])

		if line[0] == 76 {
			current = rotate(current, -distance)
		}

		if line[0] == 82 {
			current = rotate(current, distance)
		}

		if current == 0 {
			ans++
		}
	}
	return ans
}

func (s Solution) PartTwo(input any) any {
	rotations := input.([]string)

	ans, current := 0, 50

	for _, line := range rotations {
		arr := strings.FieldsFunc(line, func(r rune) bool {
			return r == 'L' || r == 'R'
		})
		distance, _ := strconv.Atoi(arr[0])

		if line[0] == 76 {
			passed := rotationsPassedZero(current, -distance)
			ans += passed
			current = rotate(current, -distance)
		}

		if line[0] == 82 {
			passed := rotationsPassedZero(current, distance)
			ans += passed
			current = rotate(current, distance)
		}
	}
	return ans
}

func GetSolution() Solution {
	return Solution{}
}

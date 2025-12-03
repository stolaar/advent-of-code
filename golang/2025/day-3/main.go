package main

import (
	"strconv"
	"strings"
)

func ProcessInput(input []string) interface{} {
	result := make([][]int, len(input)-1)

	for idx, line := range input {
		if idx == len(input)-1 {
			break
		}
		result[idx] = make([]int, len(line))

		for j, r := range line {
			result[idx][j] = int(r) - 48
		}

	}
	return result
}

func PartOne(input interface{}) interface{} {
	ratings, ans := input.([][]int), 0

	for _, joltage := range ratings {
		start := 0
		highest, secondHighest := -1, -1

		for start < len(joltage) {
			if highest < 0 || joltage[start] > joltage[highest] {
				secondHighest = highest
				highest = start
				start++
				continue
			}

			if secondHighest < 0 || (joltage[start] > joltage[secondHighest]) {
				secondHighest = start
			}

			start++
		}

		if highest == len(joltage)-1 {
			highest, secondHighest = secondHighest, highest
		} else if secondHighest < highest {
			i := highest + 1
			secondHighest = -1
			for i < len(joltage) {
				if secondHighest < 0 || joltage[i] > joltage[secondHighest] {
					secondHighest = i
				}
				i++
			}
		}

		leftStr, rightStr := strconv.Itoa(joltage[highest]), strconv.Itoa(joltage[secondHighest])

		concatenated, _ := strconv.Atoi(leftStr + rightStr)
		ans += concatenated
	}
	return ans
}

func highestTwelve(joltage []int) int {
	var ans strings.Builder
	n := len(joltage)

	i := 0
	for i < n {
		if ans.Len() == 12 {
			break
		}
		m := -1

		for i < min(n-11+ans.Len(), n) {
			if i == 0 || m < 0 || joltage[i] > joltage[m] {
				m = i
			}
			i++
		}
		i = m + 1
		ans.WriteString(strconv.Itoa(joltage[m]))
	}

	res, _ := strconv.Atoi(ans.String())

	return res
}

func PartTwo(input interface{}) interface{} {
	ratings, ans := input.([][]int), 0

	for _, joltage := range ratings {
		ans += highestTwelve(joltage)
	}

	return ans
}


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
	ranges := strings.Split(input[0], ",")

	return ranges
}

func sumOfInvalids(start, end int, anySequence bool) int {
	sum := 0

	for i := start; i <= end; i++ {
		asStr := strconv.Itoa(i)
		if anySequence {
			for j := 1; j <= len(asStr)/2; j++ {
				if isMadeOfSequence(asStr[j:], asStr[:j]) {
					sum += i
					break
				}
			}
			continue
		}
		if asStr[0:len(asStr)/2] == asStr[len(asStr)/2:] {
			sum += i
		}
	}
	return sum
}

func isMadeOfSequence(s string, seq string) bool {
	c := strings.Count(s, seq)

	return c == len(s)/len(seq) && len(s)%len(seq) == 0
}

func (s Solution) PartOne(input interface{}) interface{} {
	ranges, ans := input.([]string), 0

	for _, r := range ranges {
		rarr := strings.Split(r, "-")
		start, _ := strconv.Atoi(rarr[0])
		end, _ := strconv.Atoi(rarr[1])
		ans += sumOfInvalids(start, end, false)
	}
	return ans
}

func (s Solution) PartTwo(input interface{}) interface{} {
	ranges, ans := input.([]string), 0

	for _, r := range ranges {
		rarr := strings.Split(r, "-")
		start, _ := strconv.Atoi(rarr[0])
		end, _ := strconv.Atoi(rarr[1])
		ans += sumOfInvalids(start, end, true)
	}
	return ans
}

func GetSolution() Solution {
	return Solution{}
}

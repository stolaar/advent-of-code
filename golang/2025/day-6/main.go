package solution

import (
	"strconv"
	"strings"
)

type Solution struct{}
type problem struct {
	numbers     []int
	strs        []string
	maxDigits   int
	alignedLeft bool
	operator    byte
}

func (s Solution) ReProcessInput() bool {
	return false
}

func (s Solution) ProcessInput(input []string) any {
	dst := []problem{}

	for i, line := range input {
		n := strings.Fields(line)

		for j, f := range n {
			if i == len(input)-1 {
				dst[j].operator = f[0]
				continue
			}

			num, _ := strconv.Atoi(f)
			if len(dst) <= j {
				p := problem{
					numbers:   []int{num},
					strs:      []string{},
					maxDigits: len(f),
				}
				dst = append(dst, p)
				continue
			}

			dst[j].numbers = append(dst[j].numbers, num)
			dst[j].maxDigits = max(dst[j].maxDigits, len(f))
		}

	}
	for i, line := range input {
		s := 0
		for j, col := range dst {
			if i == len(input)-1 {
				break
			}

			part := line[s:min(col.maxDigits+s, len(line))]
			if len(part) < col.maxDigits {
				part += " "
			}
			dst[j].strs = append(dst[j].strs, part)
			s += col.maxDigits + 1
		}

	}

	return dst
}

func calculate(op byte, num, result int) int {
	if op == 42 {
		if result == 0 {
			result = 1
		}
		return result * num
	}
	return result + num
}

func (s Solution) PartOne(input any) any {
	puzzle, ans := input.([]problem), 0

	for _, p := range puzzle {
		result := 0
		for _, n := range p.numbers {
			result = calculate(p.operator, n, result)
		}
		ans += result

	}
	return ans
}

func (s Solution) PartTwo(input any) any {
	puzzle, ans := input.([]problem), 0

	for _, p := range puzzle {

		nums := make([]string, p.maxDigits)
		result := 0

		for i := p.maxDigits - 1; i >= 0; i-- {
			for _, n := range p.strs {
				if n[i] != ' ' {
					nums[i] += string(n[i])
				}
			}
		}

		for _, ns := range nums {
			n, _ := strconv.Atoi(ns)
			result = calculate(p.operator, n, result)
		}

		ans += result
	}

	return ans

}

func GetSolution() Solution {
	return Solution{}
}


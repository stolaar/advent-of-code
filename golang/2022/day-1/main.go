package solution

import (
	"strconv"
)

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	result := [][]int{}

	current := []int{}

	for _, str := range input {
		if str == "" {
			result = append(result, [][]int{current}...)
			current = []int{}
			continue
		}

		num, _ := strconv.Atoi(str)

		current = append(current, num)
	}

	return result
}

func (s Solution) PartOne(input any) any {
	numsArr := input.([][]int)
	maxc := 0

	for _, nums := range numsArr {
		count := 0
		for _, num := range nums {
			count += num
		}

		if count > maxc {
			maxc = count
		}
	}
	return maxc
}

func (s Solution) PartTwo(input any) any {
	numsArr, top1, top2, top3 := input.([][]int), 0, 0, 0

	for _, nums := range numsArr {
		count := 0
		for _, num := range nums {
			count += num
		}

		if count > top1 {
			top1, top2, top3 = count, top1, top2
		} else if count > top2 {
			top2, top3 = count, top2
		} else if count > top3 {
			top3 = count
		}
	}

	return top1 + top2 + top3
}

func GetSolution() Solution {
	return Solution{}
}

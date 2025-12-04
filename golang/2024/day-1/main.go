package solution

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	list1 []int
	list2 []int
}

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	list1, list2 := make([]int, len(input)), make([]int, len(input))

	for idx, line := range input {
		if idx == len(input)-1 {
			break
		}
		locations := strings.Fields(line)

		list1[idx], _ = strconv.Atoi(locations[0])
		list2[idx], _ = strconv.Atoi(locations[1])
	}
	return &Input{
		list1: list1,
		list2: list2,
	}
}

func (s Solution) PartOne(input any) any {
	locations := input.(*Input)

	sort.Ints(locations.list1)
	sort.Ints(locations.list2)

	result := 0

	for i := 0; i < len(locations.list1); i++ {
		result += int(math.Abs(float64(locations.list2[i]) - float64(locations.list1[i])))
	}
	return result
}

func (s Solution) PartTwo(input any) any {
	locations, dict := input.(*Input), make(map[int]int)

	for i := 0; i < len(locations.list2); i++ {
		dict[locations.list2[i]] += 1
	}

	result := 0

	for _, location := range locations.list1 {
		result += location * dict[location]
	}

	return result
}

func GetSolution() Solution {
	return Solution{}
}


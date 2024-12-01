package main

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

func ProcessInput(input []string) interface{} {
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

func PartOne(input interface{}) interface{} {
	locations := input.(*Input)

	sort.Ints(locations.list1)
	sort.Ints(locations.list2)

	result := 0

	for i := 0; i < len(locations.list1); i++ {
		result += int(math.Abs(float64(locations.list2[i]) - float64(locations.list1[i])))
	}
	return result
}

func PartTwo(input interface{}) interface{} {
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


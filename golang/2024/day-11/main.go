package main

import (
	"strconv"
	"strings"
)

func ProcessInput(input []string) interface{} {
	parts := strings.Fields(strings.Join(input, "\n"))

	nums := make([]int, len(parts))

	for idx, part := range parts {
		num, _ := strconv.Atoi(part)
		nums[idx] = num
	}

	return nums
}

func blink(nums, slice1 []int, i int) ([]int, []int) {
	var current, next []int

	if i%2 == 0 {
		current = slice1
		next = nums
	} else {
		current = nums
		next = slice1
	}

	for _, num := range current {
		if num == 0 {
			next = append(next, 1)
			continue
		}
		asStr := strconv.Itoa(num)

		if len(asStr)%2 == 0 {
			leftD, _ := strconv.Atoi(asStr[:len(asStr)/2])
			next = append(next, leftD)
			rightD, _ := strconv.Atoi(asStr[len(asStr)/2:])
			next = append(next, rightD)
			continue
		}

		next = append(next, num*2024)

	}

	if i%2 == 0 {
		nums = next
		slice1 = []int{}
	} else {
		slice1 = next
		nums = []int{}
	}

	return nums, slice1
}

func PartOne(input interface{}) interface{} {
	nums := input.([]int)

	slice1 := []int{}

	for i := 1; i <= 25; i++ {
		nums, slice1 = blink(nums, slice1, i)
	}

	return max(len(slice1), len(nums))
}

type Stone struct {
	Val      int
	Straight int
	Splits   [2]*Stone
}

func PartTwo(input interface{}) interface{} {
	nums := input.([]int)

	return nums
}

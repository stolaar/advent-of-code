package main

import (
	"strconv"
	"strings"
)

type Assignment struct {
	firstPair  [2]int
	secondPair [2]int
}

func ProcessInput(input []string) interface{} {
	result := []*Assignment{}

	for _, str := range input {
		if str == "" {
			continue
		}
		assignemnts := strings.Split(str, ",")
		first, second := assignemnts[0], assignemnts[1]
		firstPair := strings.Split(first, "-")
		secondPair := strings.Split(second, "-")

		firstMin, _ := strconv.Atoi(firstPair[0])
		firstMax, _ := strconv.Atoi(firstPair[1])

		secondMin, _ := strconv.Atoi(secondPair[0])
		secondMax, _ := strconv.Atoi(secondPair[1])

		result = append(result, &Assignment{
			firstPair:  [2]int{firstMin, firstMax},
			secondPair: [2]int{secondMin, secondMax},
		})
	}
	return result
}

func PartOne(input interface{}) interface{} {
	assignemnts, count := input.([]*Assignment), 0

	for _, assignment := range assignemnts {
		firstmin, firstmax := assignment.firstPair[0], assignment.firstPair[1]
		secondmin, secondmax := assignment.secondPair[0], assignment.secondPair[1]

		if firstmin <= secondmin && firstmax >= secondmax {
			count += 1
			continue
		}

		if secondmin <= firstmin && secondmax >= firstmax {
			count += 1
		}
	}

	return count
}

func PartTwo(input interface{}) interface{} {
	assignemnts, count := input.([]*Assignment), 0

	for _, assignment := range assignemnts {
		firstmin, firstmax := assignment.firstPair[0], assignment.firstPair[1]
		secondmin, secondmax := assignment.secondPair[0], assignment.secondPair[1]

		if firstmin <= secondmin && firstmax >= secondmax {
			count += 1
			continue
		}

		if secondmin <= firstmin && secondmax >= firstmax {
			count += 1
			continue
		}

		if firstmin <= secondmin && firstmax >= secondmin {
			count += 1
			continue
		}

		if secondmax >= firstmin && firstmax >= secondmax {
			count += 1
			continue
		}
	}

	return count
}


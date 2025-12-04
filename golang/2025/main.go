package aoc_2025

import "fmt"

type Puzzle interface {
	ProcessInput(string) any
	PartOne(input any) any
	PartTwo(input any) any
}

var Problems map[int]Puzzle

type Solutions struct{}

func (s Solutions) Run(day int) {
	fmt.Println("run 2025 day", day)
	if _, ok := Problems[day]; !ok {
		panic("day not exist")
	}

	input := Problems[day].ProcessInput("ok")
	partOneSolution := Problems[day].PartOne(input)
	fmt.Println("part one sol", partOneSolution)
	partTwoSolution := Problems[day].PartTwo(input)
	fmt.Println("part two sol", partTwoSolution)
}

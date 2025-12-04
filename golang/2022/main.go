package year

import (
	"fmt"
	"os"
	"strings"
	"time"

	solution1 "github.com/stolaar/advent-of-code/2022/day-1"
	solution2 "github.com/stolaar/advent-of-code/2022/day-2"
	solution3 "github.com/stolaar/advent-of-code/2022/day-3"
	solution4 "github.com/stolaar/advent-of-code/2022/day-4"
	solution5 "github.com/stolaar/advent-of-code/2022/day-5"
	solution6 "github.com/stolaar/advent-of-code/2022/day-6"
	solution7 "github.com/stolaar/advent-of-code/2022/day-7"
)

type Puzzle interface {
	ProcessInput([]string) any
	PartOne(input any) any
	PartTwo(input any) any
}

var Problems map[int]Puzzle = make(map[int]Puzzle)

type Solutions struct{}

func (s Solutions) Run(day int) {
	if _, ok := Problems[day]; !ok {
		panic(fmt.Sprintf("day %d not implemented", day))
	}

	inputPath := fmt.Sprintf("2022/day-%d/input.txt", day)
	f, err := os.ReadFile(inputPath)
	if err != nil {
		panic(fmt.Sprintf("failed to read input: %v", err))
	}

	lines := strings.Split(string(f), "\n")

	start := time.Now()
	input := Problems[day].ProcessInput(lines)
	fmt.Printf("Input processing: %s\n", time.Since(start))

	start = time.Now()
	p1 := Problems[day].PartOne(input)
	fmt.Printf("Part 1: %v (took %s)\n", p1, time.Since(start))

	start = time.Now()
	p2 := Problems[day].PartTwo(input)
	fmt.Printf("Part 2: %v (took %s)\n", p2, time.Since(start))
}

func init() {
	Problems[1] = solution1.GetSolution()
	Problems[2] = solution2.GetSolution()
	Problems[3] = solution3.GetSolution()
	Problems[4] = solution4.GetSolution()
	Problems[5] = solution5.GetSolution()
	Problems[6] = solution6.GetSolution()
	Problems[7] = solution7.GetSolution()
}

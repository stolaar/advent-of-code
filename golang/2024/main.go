package year

import (
	"fmt"
	"os"
	"strings"
	"time"

	solution1 "github.com/stolaar/advent-of-code/2024/day-1"
	solution2 "github.com/stolaar/advent-of-code/2024/day-2"
	solution3 "github.com/stolaar/advent-of-code/2024/day-3"
	solution4 "github.com/stolaar/advent-of-code/2024/day-4"
	solution5 "github.com/stolaar/advent-of-code/2024/day-5"
	solution6 "github.com/stolaar/advent-of-code/2024/day-6"
	solution7 "github.com/stolaar/advent-of-code/2024/day-7"
	solution8 "github.com/stolaar/advent-of-code/2024/day-8"
	solution9 "github.com/stolaar/advent-of-code/2024/day-9"
	solution10 "github.com/stolaar/advent-of-code/2024/day-10"
	solution11 "github.com/stolaar/advent-of-code/2024/day-11"
	solution12 "github.com/stolaar/advent-of-code/2024/day-12"
	solution13 "github.com/stolaar/advent-of-code/2024/day-13"
	solution14 "github.com/stolaar/advent-of-code/2024/day-14"
	solution15 "github.com/stolaar/advent-of-code/2024/day-15"
	solution16 "github.com/stolaar/advent-of-code/2024/day-16"
	solution17 "github.com/stolaar/advent-of-code/2024/day-17"
	solution18 "github.com/stolaar/advent-of-code/2024/day-18"
	solution19 "github.com/stolaar/advent-of-code/2024/day-19"
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

	inputPath := fmt.Sprintf("2024/day-%d/input.txt", day)
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
	Problems[8] = solution8.GetSolution()
	Problems[9] = solution9.GetSolution()
	Problems[10] = solution10.GetSolution()
	Problems[11] = solution11.GetSolution()
	Problems[12] = solution12.GetSolution()
	Problems[13] = solution13.GetSolution()
	Problems[14] = solution14.GetSolution()
	Problems[15] = solution15.GetSolution()
	Problems[16] = solution16.GetSolution()
	Problems[17] = solution17.GetSolution()
	Problems[18] = solution18.GetSolution()
	Problems[19] = solution19.GetSolution()
}

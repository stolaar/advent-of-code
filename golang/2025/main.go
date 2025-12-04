package year

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	solution1 "github.com/stolaar/advent-of-code/2025/day-1"
	solution2 "github.com/stolaar/advent-of-code/2025/day-2"
	solution3 "github.com/stolaar/advent-of-code/2025/day-3"
	solution4 "github.com/stolaar/advent-of-code/2025/day-4"
)

type Puzzle interface {
	ProcessInput([]string) any
	PartOne(input any) any
	PartTwo(input any) any
	ReProcessInput() bool
}

var Problems map[int]Puzzle = make(map[int]Puzzle)

type Solutions struct{}

func (s Solutions) Run(day int) {
	if _, ok := Problems[day]; !ok {
		panic(fmt.Sprintf("Day %d is not implemented", day))
	}
	problem := Problems[day]

	start := time.Now()
	f, err := os.ReadFile(fmt.Sprintf("2025/day-%d/input.txt", day))

	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}

	input := problem.ProcessInput(strings.Split(string(f), "\n"))
	end := time.Since(start)
	fmt.Printf("Input exec time: %s \n", end)

	start = time.Now()

	result := Problems[day].PartOne(input)

	end = time.Since(start)
	fmt.Printf("Part one exec time: %s \n", end)
	fmt.Println("Part one result: ", result)

	if problem.ReProcessInput() {
		input = problem.ProcessInput(strings.Split(string(f), "\n"))
	}

	start = time.Now()

	result = Problems[day].PartTwo(input)

	end = time.Since(start)
	fmt.Printf("Part two exec time: %s \n", end)
	fmt.Println("Part two result: ", result)
}

func init() {
	Problems[1] = solution1.GetSolution()
	Problems[2] = solution2.GetSolution()
	Problems[3] = solution3.GetSolution()
	Problems[4] = solution4.GetSolution()
}

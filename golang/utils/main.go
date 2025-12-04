package utils

import (
	"fmt"
	"os"
	"strconv"

	aoc2022 "github.com/stolaar/advent-of-code/2022"
	aoc2024 "github.com/stolaar/advent-of-code/2024"
	aoc2025 "github.com/stolaar/advent-of-code/2025"
)

type yearSolutions interface {
	Run(day int)
}

var yearsMap map[string]yearSolutions = make(map[string]yearSolutions)

func Run(year string, i string) {
	if _, ok := yearsMap[year]; !ok {
		fmt.Printf("No solutions for year %s", year)
		return
	}

	day, _ := strconv.Atoi(i)
	fmt.Printf("Running year %s day %s \n", year, i)
	yearsMap[year].Run(day)

}

func checkError(err error, shouldPanic bool) {
	if err != nil {
		if shouldPanic {
			panic(err)
		}
		fmt.Println(err)
	}
}

func Generate(year string, day string) {
	dir := fmt.Sprintf("%s/day-%s", year, day)
	err := os.MkdirAll(dir, os.ModePerm)
	checkError(err, true)

	f, err := os.Create(fmt.Sprintf("%s/input.txt", dir))
	checkError(err, true)
	defer f.Close()

	code := `package solution

type Solution struct {}

func (s Solution) ReProcessInput() bool {
  return false
}

func (s Solution) ProcessInput(input []string) any {
  return ""
}

func (s Solution) PartOne(input any) any {
  return ""
}

func (s Solution) PartTwo(input any) any {
  return ""
}

func GetSolution() Solution {
	return Solution{}
}
  `

	err = os.WriteFile(fmt.Sprintf("%s/main.go", dir), []byte(code), 0o644)
	checkError(err, true)
}

func init() {
	yearsMap["2022"] = aoc2022.Solutions{}
	yearsMap["2024"] = aoc2024.Solutions{}
	yearsMap["2025"] = aoc2025.Solutions{}
}

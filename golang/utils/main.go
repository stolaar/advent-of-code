package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"plugin"
	"strconv"
	"strings"
	"time"

	aoc2025 "github.com/stolaar/advent-of-code/2025"
)

func buildPlugin(year string, i string) {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", fmt.Sprintf("%s/day-%s/main.so", year, i), fmt.Sprintf("%s/day-%s/main.go", year, i))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func removePlugin(year string, i string) {
	exec.Command("rm", "-rf", fmt.Sprintf("%s/day-%s/main.so", year, i)).Run()
}

type yearSolutions interface {
	Run(day int)
}

var yearsMap map[string]yearSolutions = make(map[string]yearSolutions)

func Run(year string, i string) {
	if _, ok := yearsMap[year]; ok {
		day, _ := strconv.Atoi(i)
		fmt.Printf("Running year %s day %s \n", year, i)
		yearsMap[year].Run(day)
		return
	}

	// TODO: Keep plugin approach until all previous years are refactored
	buildPlugin(year, i)

	p, err := plugin.Open(fmt.Sprintf("%s/day-%s/main.so", year, i))
	if err != nil {
		removePlugin(year, i)
		panic(err)
	}

	start := time.Now()
	f, err := os.ReadFile(fmt.Sprintf("%s/day-%s/input.txt", year, i))
	if err != nil {
		log.Fatalf("open file error: %v", err)
		removePlugin(year, i)
		return
	}

	processInput, err := p.Lookup("ProcessInput")
	if err != nil {
		removePlugin(year, i)
		panic(err)
	}

	input := processInput.(func([]string) interface{})(strings.Split(string(f), "\n"))
	end := time.Since(start)

	fmt.Printf("Input exec time: %s \n", end)

	partOne, err := p.Lookup("PartOne")
	if err != nil {
		fmt.Println("Part one not implemented")
	}

	start = time.Now()

	result := partOne.(func(interface{}) interface{})(input)

	end = time.Since(start)
	fmt.Printf("Part one exec time: %s \n", end)
	fmt.Println("Part one result: ", result)
	partTwo, err := p.Lookup("PartTwo")
	checkError(err, false)

	if err != nil {
		fmt.Println()
		return
	}

	start = time.Now()

	result = partTwo.(func(interface{}) interface{})(input)

	end = time.Since(start)
	fmt.Printf("Part two exec time: %s \n", end)
	fmt.Println("Part two result: ", result)
	removePlugin(year, i)
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
	yearsMap["2025"] = aoc2025.Solutions{}
}

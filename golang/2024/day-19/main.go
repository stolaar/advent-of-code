package main

import (
	"strings"
)

type Design struct {
	AvailablePatterns []string
	DesiredDesign     []string
}

func ProcessInput(input []string) interface{} {
	design := Design{
		DesiredDesign: make([]string, len(input)-3),
	}

	for idx, line := range input {
		if idx == len(input)-1 {
			break
		}

		if idx == 0 {
			design.AvailablePatterns = strings.Split(line, ", ")
			continue
		}
		if idx == 1 {
			continue
		}

		design.DesiredDesign[idx-2] = line
	}
	return design
}

func isPossible(design string, maxPatternLength int, patternsDict, memo map[string]bool) bool {
	if len(design) == 0 {
		return true
	}

	if _, ok := memo[design]; ok && !memo[design] {
		return false
	}

	for i := 0; i < len(design); i++ {
		var substr strings.Builder
		substr.WriteByte(design[i])

		tests := []string{}

		if _, ok := patternsDict[substr.String()]; ok {
			tests = append(tests, substr.String())
		}

		for j := i + 1; j < min(i+maxPatternLength, len(design)); j++ {
			substr.WriteByte(design[j])

			if _, ok := patternsDict[substr.String()]; ok {
				tests = append(tests, substr.String())
			}
		}

		for k := len(tests) - 1; k >= 0; k-- {
			key := design[i+len(tests[k]):]

			memo[key] = isPossible(key, maxPatternLength, patternsDict, memo)

			if memo[key] {
				return true
			}
		}
		return false
	}
	return false
}

func PartOne(input interface{}) interface{} {
	design := input.(Design)

	patternsDict, maxPatternLength, memo := make(map[string]bool), 0, map[string]bool{}

	for _, pattern := range design.AvailablePatterns {
		if len(pattern) > maxPatternLength {
			maxPatternLength = len(pattern)
		}
		patternsDict[pattern] = true
	}

	sum := 0

	for _, word := range design.DesiredDesign {
		if isPossible(word, maxPatternLength, patternsDict, memo) {
			sum += 1
		}
	}

	return sum
}

func PartTwo(input interface{}) interface{} {
	return ""
}


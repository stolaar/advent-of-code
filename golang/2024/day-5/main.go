package main

import (
	"strconv"
	"strings"
)

type PrintQueue struct {
	rules   map[int]map[int]bool
	updates [][]int
}

func ProcessInput(input []string) interface{} {
	orderingRules, rulesProcessed, updates := make(map[int]map[int]bool), false, [][]int{}

	for i := 0; i < len(input)-1; i++ {
		if input[i] == "" {
			rulesProcessed = true
			continue
		}

		if !rulesProcessed {
			parts := strings.Split(input[i], "|")

			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])

			if _, ok := orderingRules[left]; ok {
				orderingRules[left][right] = true
			} else {
				orderingRules[left] = map[int]bool{right: true}
			}

			continue
		}

		parts := strings.Split(input[i], ",")
		update := make([]int, len(parts))

		for idx, part := range parts {
			num, _ := strconv.Atoi(part)
			update[idx] = num
		}
		updates = append(updates, update)
	}

	return &PrintQueue{
		rules:   orderingRules,
		updates: updates,
	}
}

func PartOne(input interface{}) interface{} {
	printQ, sum := input.(*PrintQueue), 0

	for _, update := range printQ.updates {
		processed, valid := make(map[int]bool), true

		for _, num := range update {
			processed[num] = true

			for key := range printQ.rules[num] {
				if _, ok := processed[key]; ok {
					valid = false
					break
				}
			}

		}

		if valid {
			sum += update[len(update)/2]
		}

	}
	return sum
}

var count = 0

func fixUpdate(processed map[int]int, rules map[int]map[int]bool, update []int) (bool, []int) {
	i := 0
	count++

	for i < len(update) {
		num := update[i]

		for key := range rules[num] {
			if _, ok := processed[key]; ok && processed[key] < i {
				update[i], update[processed[key]] = update[processed[key]], update[i]
				processed[key], processed[num] = processed[num], processed[key]
				return fixUpdate(processed, rules, update)
			}
		}
		i += 1
	}

	return true, update
}

func PartTwo(input interface{}) interface{} {
	printQ, sum := input.(*PrintQueue), 0

	for _, update := range printQ.updates {
		processed, valid, i := make(map[int]int), true, 0

		for i < len(update) {
			num := update[i]
			processed[num] = i

			for key := range printQ.rules[num] {
				if _, ok := processed[key]; ok && processed[key] < i {
					valid = false
				}
			}

			i += 1
		}

		if !valid {
			_, update = fixUpdate(processed, printQ.rules, update)
			sum += update[len(update)/2]
		}
	}

	return sum
}

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Supplies struct {
	stacks       [][]string
	instructions []*Instruction
}

type Instruction struct {
	Quantity int
	From     int
	To       int
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func ProcessInput(input []string) interface{} {
	docsProcessed, docs, instructions, maxlen := false, []string{}, []*Instruction{}, 0
	for _, str := range input {
		if str == "" {
			docsProcessed = true
			continue
		}

		if !docsProcessed {
			if len(str) > maxlen {
				maxlen = len(str)
			}
			docs = append(docs, str)
			continue
		}

		strArr := strings.Split(str, " ")
		quantity, from, to := atoi(strArr[1]), atoi(strArr[3]), atoi(strArr[5])

		instructions = append(instructions, &Instruction{
			Quantity: quantity,
			From:     from - 1,
			To:       to - 1,
		})
	}

	size := (maxlen / 4) + 1
	stacks := make([][]string, size)

	for i := len(docs) - 1; i >= 0; i-- {
		line := docs[i]
		j := 0

		for i := 0; i < maxlen; i += 4 {
			if i > len(line) {
				break
			}
			if string(line[i]) == "[" {
				stacks[j] = append(stacks[j], string(line[i+1]))
			}
			j++
		}
	}

	return Supplies{
		instructions: instructions,
		stacks:       stacks,
	}
}

func PartOne(input interface{}) interface{} {
	supplies, result := input.(Supplies), ""
	stacksjson, _ := json.Marshal(supplies.stacks)
	var stacks [][]string
	json.Unmarshal(stacksjson, &stacks)

	for _, instruction := range supplies.instructions {
		quantity, from, to := instruction.Quantity, instruction.From, instruction.To

		for i := 0; i < quantity; i++ {
			top := stacks[from][len(stacks[from])-1]
			stacks[to] = append(stacks[to], top)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	for _, crates := range stacks {
		if len(crates) > 0 {
			result += crates[len(crates)-1]
		}
	}

	return result
}

func PartTwo(input interface{}) interface{} {
	supplies, result := input.(Supplies), ""
	stacksjson, _ := json.Marshal(supplies.stacks)
	var stacks [][]string
	json.Unmarshal(stacksjson, &stacks)
	fmt.Println(stacks)

	for _, instruction := range supplies.instructions {
		quantity, from, to := instruction.Quantity, instruction.From, instruction.To

		top := stacks[from][len(stacks[from])-quantity:]
		stacks[to] = append(stacks[to], top...)
		stacks[from] = stacks[from][:len(stacks[from])-quantity]
	}

	for _, crates := range stacks {
		if len(crates) > 0 {
			result += crates[len(crates)-1]
		}
	}

	return result
}

package solution

import (
	"strings"
)

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	firstp, secondp := []string{}, []string{}
	for _, str := range input {
		if str == "" {
			continue
		}
		strArr := strings.Split(str, " ")
		opponent, me := strArr[0], strArr[1]
		firstp = append(firstp, opponent)
		secondp = append(secondp, me)
	}
	return [2][]string{firstp, secondp}
}

// LOSE 0 DRAW 3 WIN 6
// ABC XYZ 1 2 3
// rock paper scissors
// x lose y draw z win
var scores = map[[2]string]int{
	{"A", "X"}: 4,
	{"A", "Y"}: 8,
	{"A", "Z"}: 3,
	{"B", "X"}: 1,
	{"B", "Y"}: 5,
	{"B", "Z"}: 9,
	{"C", "X"}: 7,
	{"C", "Y"}: 2,
	{"C", "Z"}: 6,
}

// ABC XYZ 1 2 3
// rock paper scissors
// x lose y draw z win
var scoreIndicates = map[[2]string]string{
	{"A", "X"}: "Z",
	{"A", "Y"}: "X",
	{"A", "Z"}: "Y",
	{"B", "X"}: "X",
	{"B", "Y"}: "Y",
	{"B", "Z"}: "Z",
	{"C", "X"}: "Y",
	{"C", "Y"}: "Z",
	{"C", "Z"}: "X",
}

var outcomes = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var strengh = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var ratings = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
	"Z": "B",
	"Y": "A",
	"X": "C",
}

func score(opponent string, mine string) int {
	return scores[[2]string{opponent, mine}]
}

func (s Solution) PartOne(input any) any {
	strategyGuide := input.([2][]string)

	playerOne, playerTwo, total := strategyGuide[0], strategyGuide[1], 0

	for i := 0; i < len(playerOne); i++ {
		total += score(playerOne[i], playerTwo[i])
	}
	return total
}

func (s Solution) PartTwo(input any) any {
	strategyGuide := input.([2][]string)

	playerOne, playerTwo, total := strategyGuide[0], strategyGuide[1], 0

	for i := 0; i < len(playerOne); i++ {
		total += outcomes[playerTwo[i]] + strengh[scoreIndicates[[2]string{playerOne[i], playerTwo[i]}]]
	}
	return total
}


func GetSolution() Solution {
	return Solution{}
}

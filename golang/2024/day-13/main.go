package solution

import (
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type Machine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func processLine(line string) [2]int {
	re := regexp.MustCompile("[0-9]+")

	parts := re.FindAllString(strings.Split(line, ":")[1], -1)

	xString, yString := parts[0], parts[1]

	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)

	return [2]int{x, y}
}

func (s Solution) ProcessInput(input []string) any {
	machines := []Machine{}

	for i := 0; i < len(input)-1; i += 3 {

		machine := Machine{}

		machine.ButtonA = processLine(input[i])
		machine.ButtonB = processLine(input[i+1])
		machine.Prize = processLine(input[i+2])

		machines = append(machines, machine)

		i++
	}
	return machines
}

func processMachine(machine Machine, maxPresses int) int {
	i, tx, ty := 0, machine.Prize[0], machine.Prize[1]

	ax, ay, bx, by := machine.ButtonA[0], machine.ButtonA[1], machine.ButtonB[0], machine.ButtonB[1]

	for i <= maxPresses {
		if tx%bx == 0 && ty%by == 0 && tx/bx == ty/by {
			return (i * 3) + tx/bx
		}

		tx -= ax
		ty -= ay
		i++
	}

	return 0
}

func (s Solution) PartOne(input any) any {
	machines, total := input.([]Machine), 0

	for _, machine := range machines {
		total += processMachine(machine, 100)
	}

	return total
}

func isSolvable(buttonA, buttonB [2]int) bool {
	ax, ay, bx, by := buttonA[0], buttonA[1], buttonB[0], buttonB[1]

	return ax*by-ay*bx != 0
}

func solve(buttonA, buttonB, target [2]int) int {
	var x, y int

	a, b, d, e, tx, ty := buttonA[0], buttonA[1], buttonB[0], buttonB[1], target[0], target[1]

	coef1, coef2 := e, d

	a1 := a * coef1
	d1 := b * coef2

	tx1 := tx * coef1
	ty1 := ty * coef2

	diff1 := max(a1, d1) - min(a1, d1)
	diff2 := max(tx1, ty1) - min(tx1, ty1)

	x = diff2 / diff1

	y = (target[0] - (x * a)) / buttonB[0]

	if (x*a+y*d) != tx || (x*b+y*e) != ty {
		return 0
	}

	return (x * 3) + y
}

func (s Solution) PartTwo(input any) any {
	machines, total := input.([]Machine), 0

	for _, machine := range machines {
		machine.Prize[0] += 10000000000000
		machine.Prize[1] += 10000000000000

		total += solve(machine.ButtonA, machine.ButtonB, machine.Prize)
	}

	return total
}


func GetSolution() Solution {
	return Solution{}
}

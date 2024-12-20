package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Computer struct {
	RegisterA, RegisterB, RegisterC int
	Program                         []int
}

func generateComputer(lines []string) *Computer {
	re := regexp.MustCompile("[0-9]+")
	computer := &Computer{}

	for idx, line := range lines {
		if idx < 3 {
			registerString := re.FindString(line)
			value, _ := strconv.Atoi(registerString)

			if idx == 0 {
				computer.RegisterA = value
				continue
			}
			if idx == 1 {
				computer.RegisterB = value
				continue
			}
			computer.RegisterC = value
			continue
		}

		if idx == 3 {
			continue
		}

		instructions := re.FindAllString(line, -1)

		for _, instruction := range instructions {
			code, _ := strconv.Atoi(instruction)

			computer.Program = append(computer.Program, code)
		}
		break
	}
	return computer
}

func ProcessInput(input []string) interface{} {
	return input
}

func (c *Computer) getComboOperand(value int) int {
	if value <= 3 || value == 7 {
		return value
	}

	if value == 4 {
		return c.RegisterA
	}
	if value == 5 {
		return c.RegisterB
	}
	return c.RegisterC
}

func (c *Computer) adv(operand int) {
	c.RegisterA = int(float64(c.RegisterA) / math.Pow(2, float64(c.getComboOperand(operand))))
}

func (c *Computer) bxl(operand int) {
	c.RegisterB ^= operand
}

func (c *Computer) bst(operand int) {
	c.RegisterB = (c.getComboOperand(operand) % 8) & 7
}

func (c *Computer) jnz(operand, pointer int) int {
	if c.RegisterA == 0 {
		return pointer + 2
	}

	return operand
}

func (c *Computer) bxc() {
	c.RegisterB ^= c.RegisterC
}

func (c *Computer) out(operand int) int {
	return c.getComboOperand(operand) % 8
}

func (c *Computer) bdv(operand int) {
	c.RegisterB = int(float64(c.RegisterA) / math.Pow(2, float64(c.getComboOperand(operand))))
}

func (c *Computer) cdv(operand int) {
	c.RegisterC = int(float64(c.RegisterA) / math.Pow(2, float64(c.getComboOperand(operand))))
}

func PartOne(input interface{}) interface{} {
	computer := generateComputer(input.([]string))
	i := 0
	output := []int{}

	for i < len(computer.Program)-1 {
		opcode, operand := computer.Program[i], computer.Program[i+1]

		switch opcode {
		case 0:
			computer.adv(operand)
		case 1:
			computer.bxl(operand)
		case 2:
			computer.bst(operand)
		case 3:
			i = computer.jnz(operand, i)
			continue
		case 4:
			computer.bxc()
		case 5:
			output = append(output, computer.out(operand))
		case 6:
			computer.bdv(operand)
		case 7:
			computer.cdv(operand)
		}

		i += 2
	}

	var result strings.Builder
	fmt.Println(computer)

	for idx, out := range output {
		result.WriteString(fmt.Sprintf("%d", out))
		if idx < len(output)-1 {
			result.WriteString(",")
		}
	}

	return result.String()
}

func PartTwo(input interface{}) interface{} {
	return ""
}


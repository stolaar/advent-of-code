package main

import (
	"bytes"
	"strconv"
	"strings"
)

func ProcessInput(input []string) interface{} {
	instructions := strings.Join(input, "\n")
	return instructions
}

func PartOne(input interface{}) interface{} {
	instructions := input.(string)

	mul, leftOp, rightOp := bytes.NewBufferString(""), bytes.NewBufferString(""), bytes.NewBufferString("")

	var product int
	result := 0

	for _, r := range instructions {
		if r == 'm' || r == 'u' || r == 'l' {
			mul.WriteString(string(r))
			continue
		}

		cmd := mul.String()

		if r == '(' {
			if cmd == "mul" {
				continue
			}
		}

		if r >= 48 && r <= 57 && cmd == "mul" {
			if product == 0 {
				leftOp.WriteString(string(r))
				continue
			}

			rightOp.WriteString(string(r))
			continue
		}

		if r == ',' && cmd == "mul" && leftOp.Len() > 0 {
			leftNum, _ := strconv.Atoi(leftOp.String())
			product = leftNum
			continue
		}

		if r == ')' && cmd == "mul" && leftOp.Len() > 0 && rightOp.Len() > 0 {
			rightNum, _ := strconv.Atoi(rightOp.String())
			result += product * rightNum

			product = 0
			leftOp.Reset()
			rightOp.Reset()
		}

		product = 0
		mul.Reset()
		leftOp.Reset()
		rightOp.Reset()
		continue
	}

	return result
}

func PartTwo(input interface{}) interface{} {
	instructions := input.(string)

	mul, leftOp, rightOp := bytes.NewBufferString(""), bytes.NewBufferString(""), bytes.NewBufferString("")

	var product int
	result := 0
	mulEnabled := true
	stopIns := bytes.NewBufferString("")

	for i := 0; i < len(instructions); i++ {
		r := instructions[i]

		if r == 'd' {
			product = 0
			mul.Reset()
			leftOp.Reset()
			rightOp.Reset()

			for k := i; k < i+4; k++ {
				stopIns.WriteString(string(instructions[k]))
			}

			if stopIns.String() == "do()" {
				mulEnabled = true
				stopIns.Reset()
				i += 3
				continue
			}
			stopIns.Reset()

			for k := i; k < i+7; k++ {
				stopIns.WriteString(string(instructions[k]))
			}

			if stopIns.String() == "don't()" {
				stopIns.Reset()
				mulEnabled = false
				i += 5
				continue
			}
		}

		if !mulEnabled {
			continue
		}

		if r == 'm' || r == 'u' || r == 'l' {
			mul.WriteString(string(r))
			continue
		}

		cmd := mul.String()

		if r == '(' {
			if cmd == "mul" {
				continue
			}
		}

		if r >= 48 && r <= 57 && cmd == "mul" {
			if product == 0 {
				leftOp.WriteString(string(r))
				continue
			}

			rightOp.WriteString(string(r))
			continue
		}

		if r == ',' && cmd == "mul" && leftOp.Len() > 0 {
			leftNum, _ := strconv.Atoi(leftOp.String())
			product = leftNum
			continue
		}

		if r == ')' && cmd == "mul" && leftOp.Len() > 0 && rightOp.Len() > 0 {
			rightNum, _ := strconv.Atoi(rightOp.String())
			result += product * rightNum

			product = 0
			leftOp.Reset()
			rightOp.Reset()
		}

		product = 0
		mul.Reset()
		leftOp.Reset()
		rightOp.Reset()
		continue
	}

	return result
}

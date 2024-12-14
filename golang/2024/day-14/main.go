package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	startingPosition [2]int
	position         [2]int
	velocity         [2]int
}

func lineToRobot(line string) *Robot {
	robot := &Robot{}

	parts := strings.Split(line, " ")
	re := regexp.MustCompile("[-0-9]+")
	positions := re.FindAllString(parts[0], -1)
	velocities := re.FindAllString(parts[1], -1)

	x, _ := strconv.Atoi(positions[0])
	y, _ := strconv.Atoi(positions[1])

	robot.position = [2]int{x, y}

	x, _ = strconv.Atoi(velocities[0])
	y, _ = strconv.Atoi(velocities[1])

	robot.velocity = [2]int{x, y}

	return robot
}

func ProcessInput(input []string) interface{} {
	robots := make([]*Robot, len(input)-1)

	for i := 0; i < len(input)-1; i++ {
		robots[i] = lineToRobot(input[i])
	}
	return robots
}

func process(robot *Robot, depth, maxDepth, w, h int, memo map[*Robot]map[int][2]int) {
	memo[robot][depth] = robot.position
	if depth > maxDepth {
		return
	}

	if robot.position[0]+robot.velocity[0] >= 0 && robot.position[0]+robot.velocity[0] < w {
		robot.position[0] += robot.velocity[0]
	} else {
		if robot.position[0]+robot.velocity[0] < 0 {
			robot.position[0] = w - int(math.Abs(float64(robot.velocity[0]+robot.position[0])))
		} else {
			robot.position[0] = (robot.velocity[0] + robot.position[0]) - w
		}
	}

	if robot.position[1]+robot.velocity[1] >= 0 && robot.position[1]+robot.velocity[1] < h {
		robot.position[1] += robot.velocity[1]
	} else {
		if robot.position[1]+robot.velocity[1] < 0 {
			robot.position[1] = h - int(math.Abs(float64(robot.velocity[1]+robot.position[1])))
		} else {
			robot.position[1] = (robot.position[1] + robot.velocity[1]) - h
		}
	}

	process(robot, depth+1, maxDepth, w, h, memo)
	return
}

func robotQuadrant(x, y, w, h int) int {
	middle := x == w/2 || y == h/2

	if middle {
		return -1
	}

	left, top := x < w/2, y < h/2

	if left && top {
		return 1
	}

	if !left && top {
		return 2
	}

	if left && !top {
		return 3
	}

	return 4
}

func PartOne(input interface{}) interface{} {
	robots, memo := input.([]*Robot), map[*Robot]map[int][2]int{}

	quadrants, w, h := make([]int, 4), 101, 103

	for _, robot := range robots {
		memo[robot] = map[int][2]int{0: {robot.position[0], robot.position[1]}}

		process(robot, 0, 100, w, h, memo)

		quadrant := robotQuadrant(memo[robot][100][0], memo[robot][100][1], w, h)

		if quadrant > -1 {
			quadrants[quadrant-1]++
		}
	}

	total := 1

	for _, quadrant := range quadrants {
		total *= quadrant
	}

	return total
}

func PartTwo(input interface{}) interface{} {
	return ""
}


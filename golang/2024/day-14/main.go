package solution

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type Robot struct {
	startingPosition [2]int
	position         [2]int
	velocity         [2]int
}

func linesToRobot(lines []string) []*Robot {
	robots := make([]*Robot, len(lines)-1)

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]

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
		robots[i] = robot
	}

	return robots
}

func (s Solution) ProcessInput(input []string) any {
	return input
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

func (s Solution) PartOne(input any) any {
	robots, memo := linesToRobot(input.([]string)), map[*Robot]map[int][2]int{}

	quadrants, w, h := make([]int, 4), 101, 103

	fmt.Println(len(robots))

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

func checkTree(startRow, startCol int, grid [][]int) bool {
	baseWidth, col := 41, startCol-1
	for i := startRow - 1; i > startRow-22; i-- {
		for j := col; j < baseWidth; j++ {
			if grid[i][j] <= 0 {
				return false
			}
		}
		baseWidth -= 2
		col += 1
	}
	return true
}

func isTree(grid [][]int) bool {
	length := 43

	baseStart := -1

	lenghtFound := 0

	for i := len(grid) - 1; i >= 22 && length > 0; i-- {
		for j, col := range grid[i] {
			if col > 0 {
				lenghtFound++
			} else {
				lenghtFound = 0
			}

			if lenghtFound == length {
				baseStart = j - length
				if checkTree(i, j, grid) {
					return true
				}
				lenghtFound = 0
				continue
			}

			if len(grid[i])-j+lenghtFound < length {
				break
			}
		}
		if baseStart > -1 {
			break
		}
	}

	return false
}

func (s Solution) PartTwo(input any) any {
	robots, i := linesToRobot(input.([]string)), 1

	w, h := 101, 103

	grid := make([][]int, h)

	for i := range grid {
		grid[i] = make([]int, w)
	}

	for _, robot := range robots {
		x, y := robot.position[0], robot.position[1]
		grid[y][x]++
	}

	// for i < math.MaxInt64 {
	// 	for _, robot := range robots {
	// 		px, py := robot.position[0], robot.position[1]
	// 		if _, ok := memo[robot]; !ok {
	// 			memo[robot] = map[int][2]int{}
	// 		}
	// 		process(robot, 0, math.MaxInt64, w, h, memo)
	//
	// 		x, y := robot.position[0], robot.position[1]
	// 		grid[py][px]--
	// 		grid[y][x]--
	// 	}
	// 	if isTree(grid) {
	// 		return i
	// 	}
	//
	// 	i++
	// }
	return i
}

func GetSolution() Solution {
	return Solution{}
}

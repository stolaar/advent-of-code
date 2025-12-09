package solution

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}

type point struct {
	x, y, z         int
	pair            *point
	closestDistance int
}

type circuit struct {
	points []*point
	size   int
}

func distance(p1, p2 point) float64 {
	dx := float64(p1.x) - float64(p2.x)
	dy := float64(p1.y) - float64(p2.y)
	dz := float64(p1.z) - float64(p2.z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (s Solution) ReProcessInput() bool {
	return true
}

func (s Solution) ProcessInput(input []string) any {
	p := make([]*point, len(input))

	for i, line := range input {
		points := strings.Split(line, ",")
		x, _ := strconv.Atoi(points[0])
		y, _ := strconv.Atoi(points[1])
		z, _ := strconv.Atoi(points[2])

		pp := &point{
			x, y, z,
			nil,
			math.MaxInt,
		}
		p[i] = pp
	}
	return p
}

func (s Solution) PartOne(input any) any {
	puzzle := input.([]*point)
	circuits := []*circuit{}

	pointsMap := map[*point]*circuit{}
	distances := [][3]any{}

	for i := range puzzle {
		for j := i + 1; j < len(puzzle); j++ {
			d := distance(*puzzle[i], *puzzle[j])

			distances = append(distances, [3]any{d, puzzle[i], puzzle[j]})
		}
	}

	slices.SortFunc(distances, func(d1, d2 [3]any) int {
		return int(d1[0].(float64) - d2[0].(float64))
	})

	for i := range 1000 {
		distance := distances[i]
		p1, p2 := distance[1].(*point), distance[2].(*point)

		pm1, ok1 := pointsMap[p1]
		pm2, ok2 := pointsMap[p2]

		if ok1 && ok2 {
			if pm1 == pm2 {
				continue
			}

			pm1.size += pm2.size

			for _, p := range pm2.points {
				pointsMap[p] = pm1
				pm1.points = append(pm1.points, p)
			}

			pm2.size = 0
			pointsMap[p2] = pm1
			continue
		}

		if !ok1 && ok2 {
			pm2.size += 1
			pm2.points = append(pm2.points, p1)
			pointsMap[p1] = pm2
			continue
		}

		if ok1 && !ok2 {
			pm1.size += 1
			pm1.points = append(pm1.points, p2)
			pointsMap[p2] = pm1
			continue
		}

		c := &circuit{
			points: []*point{p1, p2},
			size:   2,
		}

		circuits = append(circuits, c)
		pointsMap[p1] = c
		pointsMap[p2] = c
	}

	slices.SortFunc(circuits, func(c1, c2 *circuit) int {
		return c2.size - c1.size
	})

	ans := 1
	for i := range 3 {
		ans *= circuits[i].size
	}

	return ans
}

func (s Solution) PartTwo(input any) any {
	puzzle := input.([]*point)
	circuits := []*circuit{}

	pointsMap := map[*point]*circuit{}
	distances := [][3]any{}

	for i := range puzzle {
		for j := i + 1; j < len(puzzle); j++ {
			d := distance(*puzzle[i], *puzzle[j])

			distances = append(distances, [3]any{d, puzzle[i], puzzle[j]})
		}
	}

	slices.SortFunc(distances, func(d1, d2 [3]any) int {
		return int(d1[0].(float64) - d2[0].(float64))
	})

	for _, distance := range distances {
		p1, p2 := distance[1].(*point), distance[2].(*point)

		pm1, ok1 := pointsMap[p1]
		pm2, ok2 := pointsMap[p2]

		if ok1 && ok2 {
			if pm1 == pm2 {
				continue
			}

			pm1.size += pm2.size

			if pm1.size == len(puzzle) {
				return p1.x * p2.x
			}

			for _, p := range pm2.points {
				pointsMap[p] = pm1
				pm1.points = append(pm1.points, p)
			}

			pm2.size = 0
			pointsMap[p2] = pm1
			continue
		}

		if !ok1 && ok2 {
			pm2.size += 1
			pm2.points = append(pm2.points, p1)
			pointsMap[p1] = pm2
			if pm2.size == len(puzzle) {
				return p1.x * p2.x
			}
			continue
		}

		if ok1 && !ok2 {
			pm1.size += 1
			pm1.points = append(pm1.points, p2)
			pointsMap[p2] = pm1
			if pm1.size == len(puzzle) {
				return p1.x * p2.x
			}
			continue
		}

		c := &circuit{
			points: []*point{p1, p2},
			size:   2,
		}

		circuits = append(circuits, c)
		pointsMap[p1] = c
		pointsMap[p2] = c
	}

	return 0
}

func GetSolution() Solution {
	return Solution{}
}

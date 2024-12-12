package main

type Plot struct {
	Val                                                  byte
	Y, X                                                 int
	Visited, SidesCounted                                bool
	LeftFence, TopFence, RightFence, BottomFence         bool
	LeftVisited, TopVisited, RightVisited, BottomVisited bool
}

type Direction int

const (
	Up Direction = iota
	Down
	Right
	Left
)

type Garden struct {
	grid      [][]*Plot
	plotAreas map[[2]int]*[2]int
	plotSides map[[2]int]int
}

func (g *Garden) addFence(key [2]int, plot *Plot) {
	plotFences := 4
	plot.Visited = true

	plot.TopFence = true
	plot.LeftFence = true
	plot.RightFence = true
	plot.BottomFence = true

	var right, down, left, up *Plot

	if plot.X < len(g.grid[0])-1 {
		if g.grid[plot.Y][plot.X+1].Val == plot.Val {
			right = g.grid[plot.Y][plot.X+1]
			plotFences--
			plot.RightFence = false
		}
	}

	if plot.Y < len(g.grid)-1 {
		if g.grid[plot.Y+1][plot.X].Val == plot.Val {
			down = g.grid[plot.Y+1][plot.X]
			plotFences--
			plot.BottomFence = false
		}
	}

	if plot.X > 0 {
		if g.grid[plot.Y][plot.X-1].Val == plot.Val {
			left = g.grid[plot.Y][plot.X-1]
			plotFences--
			plot.LeftFence = false
		}
	}

	if plot.Y > 0 {
		if g.grid[plot.Y-1][plot.X].Val == plot.Val {
			up = g.grid[plot.Y-1][plot.X]
			plotFences--
			plot.TopFence = false
		}
	}

	if _, ok := g.plotAreas[key]; ok {
		g.plotAreas[key][0] += plotFences
		g.plotAreas[key][1]++
	} else {
		g.plotAreas[key] = &[2]int{plotFences, 1}
	}

	if up != nil && !up.Visited {
		g.addFence(key, up)
	}

	if down != nil && !down.Visited {
		g.addFence(key, down)
	}

	if right != nil && !right.Visited {
		g.addFence(key, right)
	}

	if left != nil && !left.Visited {
		g.addFence(key, left)
	}
}

func (g *Garden) checkLeftSides(key [2]int, plot *Plot) {
	if !plot.LeftFence || plot.LeftVisited {
		plot.LeftVisited = true
		return
	}
	plot.LeftVisited = true
	g.plotSides[key]++

	y := plot.Y - 1

	for y >= 0 && g.grid[y][plot.X].Val == plot.Val {
		if g.grid[y][plot.X].LeftFence {
			g.grid[y][plot.X].LeftVisited = true
			y--
			continue
		}
		break
	}

	y = plot.Y + 1

	for y < len(g.grid) && g.grid[y][plot.X].Val == plot.Val {
		if g.grid[y][plot.X].LeftFence {
			g.grid[y][plot.X].LeftVisited = true
			y++
			continue
		}
		break
	}
}

func (g *Garden) checkRightSides(key [2]int, plot *Plot) {
	if !plot.RightFence || plot.RightVisited {
		plot.RightVisited = true
		return
	}

	plot.RightVisited = true
	g.plotSides[key]++

	y := plot.Y - 1

	for y >= 0 && g.grid[y][plot.X].Val == plot.Val {
		if g.grid[y][plot.X].RightFence {
			g.grid[y][plot.X].RightVisited = true
			y--
			continue
		}
		break
	}

	y = plot.Y + 1

	for y < len(g.grid) && g.grid[y][plot.X].Val == plot.Val {
		if g.grid[y][plot.X].RightFence {
			g.grid[y][plot.X].RightVisited = true
			y++
			continue
		}
		break
	}
}

func (g *Garden) checkTopSides(key [2]int, plot *Plot) {
	if !plot.TopFence || plot.TopVisited {
		plot.TopVisited = true
		return
	}
	plot.TopVisited = true
	g.plotSides[key]++

	x := plot.X - 1

	for x >= 0 && g.grid[plot.Y][x].Val == plot.Val {
		if g.grid[plot.Y][x].TopFence {
			g.grid[plot.Y][x].TopVisited = true
			x--
			continue
		}
		g.grid[plot.Y][x].TopVisited = true
		break
	}

	x = plot.X + 1

	for x < len(g.grid[0]) && g.grid[plot.Y][x].Val == plot.Val {
		if g.grid[plot.Y][x].TopFence {
			g.grid[plot.Y][x].TopVisited = true
			x++
			continue
		}
		g.grid[plot.Y][x].TopVisited = true
		break
	}
}

func (g *Garden) checkDownSides(key [2]int, plot *Plot) {
	if !plot.BottomFence || plot.BottomVisited {
		plot.BottomVisited = true
		return
	}
	plot.BottomVisited = true
	g.plotSides[key]++

	x := plot.X - 1

	for x >= 0 && g.grid[plot.Y][x].Val == plot.Val {
		if g.grid[plot.Y][x].BottomFence {
			g.grid[plot.Y][x].BottomVisited = true
			x--
			continue
		}
		g.grid[plot.Y][x].BottomVisited = true
		break
	}

	x = plot.X + 1

	for x < len(g.grid[0]) && g.grid[plot.Y][x].Val == plot.Val {
		if g.grid[plot.Y][x].BottomFence {
			g.grid[plot.Y][x].BottomVisited = true
			x++
			continue
		}
		g.grid[plot.Y][x].BottomVisited = true
		break
	}
}

func (g *Garden) calcSides(key [2]int, plot *Plot) {
	plot.SidesCounted = true

	var right, down, left, up *Plot

	if plot.X < len(g.grid[0])-1 {
		if g.grid[plot.Y][plot.X+1].Val == plot.Val {
			right = g.grid[plot.Y][plot.X+1]
		}
	}

	if plot.Y < len(g.grid)-1 {
		if g.grid[plot.Y+1][plot.X].Val == plot.Val {
			down = g.grid[plot.Y+1][plot.X]
		}
	}

	if plot.X > 0 {
		if g.grid[plot.Y][plot.X-1].Val == plot.Val {
			left = g.grid[plot.Y][plot.X-1]
		}
	}

	if plot.Y > 0 {
		if g.grid[plot.Y-1][plot.X].Val == plot.Val {
			up = g.grid[plot.Y-1][plot.X]
		}
	}

	g.checkRightSides(key, plot)
	g.checkDownSides(key, plot)
	g.checkLeftSides(key, plot)
	g.checkTopSides(key, plot)

	if up != nil && !up.SidesCounted {
		g.calcSides(key, up)
	}

	if down != nil && !down.SidesCounted {
		g.calcSides(key, down)
	}

	if right != nil && !right.SidesCounted {
		g.calcSides(key, right)
	}

	if left != nil && !left.SidesCounted {
		g.calcSides(key, left)
	}
}

func stringsToGarden(str interface{}) *Garden {
	input := str.([]string)
	grid := make([][]*Plot, len(input)-1)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]*Plot, len(input[0]))

		for j := range grid[i] {
			grid[i][j] = &Plot{
				Val: input[i][j],
				X:   j,
				Y:   i,
			}
		}
	}

	return &Garden{
		grid:      grid,
		plotAreas: map[[2]int]*[2]int{},
		plotSides: map[[2]int]int{},
	}
}

func ProcessInput(input []string) interface{} {
	return input
}

func PartOne(input interface{}) interface{} {
	garden := stringsToGarden(input)

	for _, row := range garden.grid {
		for _, col := range row {
			if !col.Visited {
				garden.addFence([2]int{col.Y, col.X}, col)
			}
		}
	}

	total := 0
	for _, val := range garden.plotAreas {
		perimeter, area := val[0], val[1]

		total += perimeter * area
	}
	return total
}

func PartTwo(input interface{}) interface{} {
	garden := stringsToGarden(input)

	for _, row := range garden.grid {
		for _, col := range row {
			if !col.Visited {
				key := [2]int{col.Y, col.X}
				garden.addFence(key, col)
				garden.calcSides(key, col)
			}
		}
	}

	total := 0
	for key, val := range garden.plotAreas {
		area, sides := val[1], garden.plotSides[key]

		total += sides * area
	}
	return total
}

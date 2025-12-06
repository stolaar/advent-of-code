package solution

type Solution struct{}

func (s Solution) ReProcessInput() bool {
	return false
}

type grid [][]*cell

type cell struct {
	rollpaper bool
	x, y      int
	removed   bool
}

func (s Solution) ProcessInput(input []string) any {
	puzzle := make(grid, len(input)-1)

	for i, line := range input {
		row := make([]*cell, len(line))

		for j, c := range line {
			row[j] = &cell{
				rollpaper: c == 64,
				x:         j,
				y:         i,
			}
		}
		puzzle[i] = row
	}
	return puzzle
}

func (g grid) maxX() int {
	return len(g[0]) - 1
}

func (g grid) maxY() int {
	return len(g) - 1
}

func (g grid) canBeForklifted(c *cell) bool {
	numberOfPossibleLifts := 0

	if c.x > 0 && g[c.y][c.x-1].rollpaper && !g[c.y][c.x-1].removed {
		numberOfPossibleLifts++
	}

	if c.x < g.maxX() && g[c.y][c.x+1].rollpaper && !g[c.y][c.x+1].removed {
		numberOfPossibleLifts++
	}

	if c.y > 0 && g[c.y-1][c.x].rollpaper && !g[c.y-1][c.x].removed {
		numberOfPossibleLifts++
	}

	if c.y < g.maxY() && g[c.y+1][c.x].rollpaper && !g[c.y+1][c.x].removed {
		numberOfPossibleLifts++
	}

	if c.x > 0 && c.y > 0 && g[c.y-1][c.x-1].rollpaper && !g[c.y-1][c.x-1].removed {
		numberOfPossibleLifts++
	}

	if c.x < g.maxX() && c.y > 0 && g[c.y-1][c.x+1].rollpaper && !g[c.y-1][c.x+1].removed {
		numberOfPossibleLifts++
	}

	if c.y < g.maxY() && c.x > 0 && g[c.y+1][c.x-1].rollpaper && !g[c.y+1][c.x-1].removed {
		numberOfPossibleLifts++
	}

	if c.y < g.maxY() && c.x < g.maxX() && g[c.y+1][c.x+1].rollpaper && !g[c.y+1][c.x+1].removed {
		numberOfPossibleLifts++
	}
	return numberOfPossibleLifts < 4
}

func (s Solution) PartOne(input any) any {
	puzzle, ans := input.(grid), 0

	for _, row := range puzzle {
		for _, col := range row {
			if col.rollpaper && puzzle.canBeForklifted(col) {
				ans++
			}
		}
	}

	return ans
}

func (s Solution) PartTwo(input any) any {
	puzzle, ans := input.(grid), 0

	run := true
	for run {
		run = false
		for _, row := range puzzle {
			for _, col := range row {
				if col.rollpaper && !col.removed && puzzle.canBeForklifted(col) {
					col.removed = true
					run = true
					ans++
				}
			}
		}
	}

	return ans
}

func GetSolution() Solution {
	return Solution{}
}

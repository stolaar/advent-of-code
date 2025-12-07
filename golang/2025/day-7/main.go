package solution

type Solution struct{}

type puzzleInput struct {
	grid  grid
	start *cell
}

type cell struct {
	x, y                    int
	isSplitter, beamPassing bool
	timesSplit              int
	passed                  bool
	possiblePathsToHere     int
}

type grid [][]*cell

func (s Solution) ReProcessInput() bool {
	return true
}

func (s Solution) ProcessInput(input []string) any {
	dst := puzzleInput{}

	g := make(grid, len(input))
	for i, r := range input {
		row := make([]*cell, len(r))
		for j, c := range r {
			cl := &cell{
				x:          j,
				y:          i,
				isSplitter: c == 94,
			}

			if c == 83 {
				dst.start = cl
			}
			row[j] = cl
		}
		g[i] = row
	}
	dst.grid = g
	return dst
}

func (p puzzleInput) streamBeam(col *cell) int {
	if col.y+1 >= len(p.grid) || col.passed {
		if col.beamPassing && col.y+1 >= len(p.grid) {
			return col.possiblePathsToHere + 1
		}
		return col.possiblePathsToHere
	}

	col.passed = true

	downCol := p.grid[col.y+1][col.x]
	if downCol.isSplitter {
		if col.beamPassing && ((downCol.x-1 > 0 && !p.grid[downCol.y][downCol.x-1].isSplitter) || (downCol.x+1 < len(p.grid[0]) && !p.grid[downCol.y][downCol.x+1].isSplitter)) {
			col.timesSplit++
		}
		if downCol.x-1 > 0 {
			res := p.streamBeam(p.grid[downCol.y][downCol.x-1])
			col.possiblePathsToHere += res
		}
		if downCol.x+1 < len(p.grid[0]) {
			res := p.streamBeam(p.grid[downCol.y][downCol.x+1])
			col.possiblePathsToHere += res
		}
	} else {
		downCol.beamPassing = true
		res := p.streamBeam(downCol)
		col.possiblePathsToHere += res
	}

	return col.possiblePathsToHere
}

func (s Solution) PartOne(input any) any {
	puzzle := input.(puzzleInput)

	puzzle.streamBeam(puzzle.start)

	ans := 0

	for i := puzzle.start.y; i < len(puzzle.grid); i++ {
		for _, col := range puzzle.grid[i] {
			if col.timesSplit > 0 {
				ans++
			}
		}
	}
	return ans
}

func (s Solution) PartTwo(input any) any {
	puzzle := input.(puzzleInput)

	puzzle.streamBeam(puzzle.start)
	return puzzle.start.possiblePathsToHere + 1
}

func GetSolution() Solution {
	return Solution{}
}


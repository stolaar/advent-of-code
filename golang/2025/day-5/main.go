package solution

import (
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}
type ingredientRange struct {
	start, end int
}
type puzzleInput struct {
	ingredientIdRanges   []ingredientRange
	availableIngredients []int
}

func (s Solution) ReProcessInput() bool {
	return false
}

func (s Solution) ProcessInput(input []string) any {
	p := puzzleInput{
		ingredientIdRanges:   []ingredientRange{},
		availableIngredients: []int{},
	}
	i := 0

	for i = 0; i < len(input); i++ {
		if input[i] == "" {
			i++
			break
		}

		rng := strings.Split(input[i], "-")

		start, _ := strconv.Atoi(rng[0])
		end, _ := strconv.Atoi(rng[1])
		ingRange := ingredientRange{
			start: start,
			end:   end,
		}

		p.ingredientIdRanges = append(p.ingredientIdRanges, ingRange)
	}

	for i < len(input) {
		id, _ := strconv.Atoi(input[i])
		p.availableIngredients = append(p.availableIngredients, id)
		i++
	}

	return p
}

func (s Solution) PartOne(input any) any {
	puzzle := input.(puzzleInput)

	slices.SortFunc(puzzle.ingredientIdRanges, func(a, b ingredientRange) int {
		if a.start == b.start {
			return a.end - b.end
		}

		return a.start - b.start
	})
	ans := 0

	for _, available := range puzzle.availableIngredients {

		for _, r := range puzzle.ingredientIdRanges {
			if available >= r.start && available <= r.end {
				ans++
				break
			}

			if available < r.start {
				break
			}
		}
	}
	return ans
}

func (s Solution) PartTwo(input any) any {
	puzzle := input.(puzzleInput)

	slices.SortFunc(puzzle.ingredientIdRanges, func(a, b ingredientRange) int {
		if a.start == b.start {
			return a.end - b.end
		}

		return a.start - b.start
	})

	merged := []ingredientRange{}

	ans, i := 0, 0
	for i < len(puzzle.ingredientIdRanges) {
		sr := puzzle.ingredientIdRanges[i]

		j := i + 1
		for j = i + 1; j < len(puzzle.ingredientIdRanges); j++ {
			nr := puzzle.ingredientIdRanges[j]

			if nr.start <= sr.end {
				sr.end = max(sr.end, nr.end)
				continue
			}
			break
		}
		merged = append(merged, sr)
		i = j
	}

	for _, r := range merged {
		ans += r.end - r.start + 1
	}

	return ans
}

func GetSolution() Solution {
	return Solution{}
}

package solution

type Solution struct{}

func (s Solution) ProcessInput(input []string) any {
	return input
}

func (s Solution) PartOne(input any) any {
	buffer := input.([]string)[0]

	l, r := 0, 1
	for r < len(buffer) {
		if r-l == 4 {
			return r
		}
		i := l

		ok := true
		for i < r {
			if buffer[i] == buffer[r] {
				l = i + 1
				if r == l {
					r += 1
				}
				ok = false
				break
			}
			i += 1
		}
		if ok {
			r += 1
		}
	}

	return r
}

func (s Solution) PartTwo(input any) any {
	buffer := input.([]string)[0]

	l, r := 0, 1
	for r < len(buffer) {
		if r-l == 14 {
			return r
		}
		i := l

		ok := true
		for i < r {
			if buffer[i] == buffer[r] {
				l = i + 1
				if r == l {
					r += 1
				}
				ok = false
				break
			}
			i += 1
		}
		if ok {
			r += 1
		}
	}

	return r
}

func GetSolution() Solution {
	return Solution{}
}

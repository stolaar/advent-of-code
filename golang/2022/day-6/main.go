package main

func ProcessInput(input []string) interface{} {
	return input
}

func PartOne(input interface{}) interface{} {
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

func PartTwo(input interface{}) interface{} {
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

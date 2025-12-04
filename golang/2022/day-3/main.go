package solution

type Solution struct{}

type Rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func (s Solution) ProcessInput(input []string) any {
	result := []*Rucksack{}

	for _, str := range input {
		first, second := str[:len(str)/2], str[len(str)/2:]

		result = append(result, &Rucksack{
			firstCompartment:  first,
			secondCompartment: second,
		})
	}

	return result
}

func getPriority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r) - 96
	}
	return int(r) - 38
}

func (s Solution) PartOne(input any) any {
	rucksacks, result := input.([]*Rucksack), 0

	for _, rucksack := range rucksacks {
		m := make(map[rune]bool, len(rucksack.firstCompartment))
		for _, r := range rucksack.firstCompartment {
			m[r] = true
		}

		for _, r := range rucksack.secondCompartment {
			if m[r] {
				result += getPriority(r)
				break
			}
		}
	}
	return result
}

func (s Solution) PartTwo(input any) any {
	rucksacks, result := input.([]*Rucksack), 0

	for i := 0; i < len(rucksacks)-3; i += 3 {
		m := make(map[rune]bool, len(rucksacks[i].firstCompartment))
		common := map[rune]bool{}

		for j := i; j < i+3; j++ {
			combined := rucksacks[j].firstCompartment + rucksacks[j].secondCompartment

			for _, r := range combined {
				if j == i {
					m[r] = true
					continue
				}
				if j == i+1 {
					if m[r] {
						common[r] = true
					}
					continue
				}

				if common[r] {

					result += getPriority(r)
					break
				}
			}
		}
	}
	return result
}


func GetSolution() Solution {
	return Solution{}
}

package solution

import (
	"strconv"
	"strings"
)

type Solution struct{}

type Equation struct {
	TestVal int
	Nums    []int
}

func (s Solution) ProcessInput(input []string) any {
	equations := make([]Equation, len(input)-1)

	for i := 0; i < len(input)-1; i++ {
		parts := strings.Split(input[i], ":")

		testVal, _ := strconv.Atoi(parts[0])

		numParts := strings.Fields(parts[1])

		nums := make([]int, len(numParts))

		for idx, part := range numParts {
			num, _ := strconv.Atoi(part)
			nums[idx] = num
		}

		eq := Equation{
			TestVal: testVal,
			Nums:    nums,
		}

		equations[i] = eq
	}
	return equations
}

func dp(target int, current int, j int, nums []int) bool {
	sum, mul := current+nums[j], current*nums[j]
	isLastIndex := j == len(nums)-1

	if (sum == target || mul == target) && isLastIndex {
		return true
	}

	if sum > target && mul > target {
		return false
	}

	if isLastIndex {
		return false
	}

	return dp(target, sum, j+1, nums) || dp(target, mul, j+1, nums)
}

func concat(i, j int) int {
	is, js := strconv.Itoa(i), strconv.Itoa(j)

	res, _ := strconv.Atoi(is + js)

	return res
}

func dp2(target int, acc int, j int, nums []int) bool {
	if j >= len(nums) {
		return false
	}

	sum, mul, merged := acc+nums[j], acc*nums[j], concat(acc, nums[j])
	isLastIndex := j == len(nums)-1

	if (sum == target || mul == target || merged == target) && isLastIndex {
		return true
	}

	if sum > target && mul > target && merged > target {
		return false
	}

	if isLastIndex {
		return false
	}

	return dp2(target, sum, j+1, nums) || dp2(target, mul, j+1, nums) || dp2(target, merged, j+1, nums)
}

func (s Solution) PartOne(input any) any {
	equations := input.([]Equation)

	sum := 0

	for _, eq := range equations {
		if dp(eq.TestVal, eq.Nums[0], 1, eq.Nums) {
			sum += eq.TestVal
		}
	}
	return sum
}

func (s Solution) PartTwo(input any) any {
	equations := input.([]Equation)

	sum := 0

	for _, eq := range equations {
		if dp2(eq.TestVal, eq.Nums[0], 1, eq.Nums) {
			sum += eq.TestVal
		}
	}

	return sum
}


func GetSolution() Solution {
	return Solution{}
}

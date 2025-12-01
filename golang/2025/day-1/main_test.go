package main

import "testing"

func TestRotationPassedZero(t *testing.T) {
	tests := []struct {
		current  int
		rotation int
		expected int
		passed   int
	}{
		{99, 3, 2, 1},
		{99, 99, 98, 1},
		{99, 100, 99, 1},
		{99, 1, 0, 1},

		{1, -2, 99, 1},
		{1, -299, 2, 3},

		{1, -499, 2, 5},
		{1, 499, 0, 5},

		{50, 1000, 50, 10},
	}

	for _, test := range tests {
		result := rotate(test.current, test.rotation)
		passed := rotationsPassedZero(test.current, test.rotation)

		if result != test.expected {
			t.Errorf("Result failed - current %d, rotate %d - Got %d - expected %d", test.current, test.rotation, result, test.expected)
		}

		if passed != test.passed {
			t.Errorf("Count failed - current %d, rotate %d - Got %d - expected %d", test.current, test.rotation, passed, test.passed)
		}
	}

}

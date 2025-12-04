package solution

import "testing"

func TestInvalidSum(t *testing.T) {
	tests := []struct {
		start, end int
		expected   int
	}{
		{1188511880, 1188511890, 1188511885},
	}

	for _, test := range tests {
		result := sumOfAnyDuplicateSeq(test.start, test.end)

		if result != test.expected {
			t.Errorf("Expected %d got %d", test.expected, result)
		}

	}
}

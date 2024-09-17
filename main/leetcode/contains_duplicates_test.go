package leetcode

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 1}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, true},
	}

	for _, test := range tests {
		result := ContainsDuplicate(test.input)
		if result != test.expected {
			t.Errorf("containsDuplicate(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

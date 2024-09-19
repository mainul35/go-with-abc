package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{1, 2, 3, 0}, []int{0, 0, 0, 6}},
		{[]int{1}, []int{1}},
	}

	for _, test := range tests {
		result := ProductExceptSelf(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ProductExceptSelf(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

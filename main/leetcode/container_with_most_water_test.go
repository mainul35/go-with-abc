package leetcode

import "testing"

// Problem link: https://leetcode.com/problems/container-with-most-water/description/

func Test_calculateMaxAreaCase1(t *testing.T) {

	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(heights)
	t.Logf("Expected: %d, got: %d", 49, res)
}

func Test_calculateMaxAreaCase2(t *testing.T) {
	heights := []int{1, 1}
	res := maxArea(heights)
	t.Logf("Expected: %d, got: %d", 1, res)
}

func Test_calculateMaxAreaCase3(t *testing.T) {
	heights := []int{1, 2, 1}
	res := maxArea(heights)
	t.Logf("Expected: %d, got: %d", 2, res)
}

func Test_calculateMaxAreaCase4(t *testing.T) {
	heights := []int{1, 2}
	res := maxArea(heights)
	t.Logf("Expected: %d, got: %d", 1, res)
}

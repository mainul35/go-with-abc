package leetcode

import "math"

func maxArea(height []int) int {
	n := len(height)
	L := 0
	R := n - 1
	max := math.MinInt

	for L < R {
		tempMax := (R - L) * int(math.Min(float64(height[L]), float64(height[R])))
		max = int(math.Max(float64(max), float64(tempMax)))
		if height[L] < height[R] {
			L++
		} else {
			R--
		}
	}
	return max
}

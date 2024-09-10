package leetcode

// Leetcode 2579
// The interval of the nth item and (n-1)th item is 4 * (n-1)
// And the total colored cells of nth item is 4*(n-1) + 4*(n-2) + 4*(n-3) + ... + 1
// Optimization technique: Memoization

func ColoredCells(n int) int64 {
	return colorCellsRecursiveBackward(n)
}

func colorCellsRecursiveBackward(n int) int64 {
	var res int64 = 0
	if n > 1 {
		n = n - 1
		res = 4 * int64(n)
		res += colorCellsRecursiveBackward(n)
	} else {
		res = 1
	}
	return res
}

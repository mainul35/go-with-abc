package leetcode

// Check if the number can be divided by 2. If yes, then the number itself is the smallest
// otherwise, multiple the number by 2 and return
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return n * 2
	}
}

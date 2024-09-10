package leetcode

// Problem 2481

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	} else if n%2 == 0 {
		return n / 2
	} else {
		return n
	}
}

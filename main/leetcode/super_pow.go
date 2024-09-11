package leetcode

// Leetcode 372

func modExp(a, b, mod int) int {
	result := 1
	a = a % mod

	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % mod
		}
		a = (a * a) % mod
		b /= 2
	}

	return result
}

func SuperPow(a int, b []int) int {
	result := 1
	for _, digit := range b {
		result = modExp(result, 10, 1337) * modExp(a, digit, 1337) % 1337
	}
	return result
}

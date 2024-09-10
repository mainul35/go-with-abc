package leetcode

func FindGCD(nums []int) int {
	min := 1000
	max := 1
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	gcd := gcd(max, min)

	return gcd
}

func gcd(max int, min int) int {
	if min == 0 {
		return max
	} else {
		return gcd(min, max%min)
	}
}

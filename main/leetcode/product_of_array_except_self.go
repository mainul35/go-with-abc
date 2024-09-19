package leetcode

func ProductExceptSelf(nums []int) []int {
	len := len(nums)
	output := make([]int, len)
	for i := 0; i < len; i++ {
		output[i] = 1
	}
	mul := 1
	for i := 0; i < len; i++ {
		output[i] *= mul
		mul *= nums[i]
	}

	mul = 1
	for i := len - 1; i > -1; i-- {
		output[i] *= mul
		mul *= nums[i]
	}
	return output
}

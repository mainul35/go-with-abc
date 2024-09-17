package leetcode

import (
	"sort"
)

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

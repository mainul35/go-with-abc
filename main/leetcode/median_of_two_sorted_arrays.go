package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums) == 1 {
		return float64(nums[0])
	}
	len := len(nums)
	halfLen := len / 2
	if len%2 == 0 {
		return float64(nums[halfLen-1]+nums[halfLen]) / 2
	} else {
		return float64(nums[halfLen])
	}
}

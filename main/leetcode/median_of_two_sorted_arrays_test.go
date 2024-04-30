package leetcode

import (
	"testing"
)

func Test_findMedianSortedArrays_withOddLength(t *testing.T) {
	nums1, nums2 := []int{1, 3}, []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("Expected %f, Got %f", 2.0, res)
}

func Test_findMedianSortedArrays_withEvenLength(t *testing.T) {
	nums1, nums2 := []int{1, 2}, []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("Expected %f, Got %f", 2.5, res)
}

func Test_findMedianSortedArrays_withOneItem(t *testing.T) {
	nums1, nums2 := []int{}, []int{1}
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("Expected %f, Got %f", 1.0, res)
}

func Test_findMedianSortedArrays_fixWrongAnswer(t *testing.T) {
	nums1, nums2 := []int{}, []int{2, 3}
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("Expected %f, Got %f", 2.5, res)
}

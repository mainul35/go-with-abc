package leetcode

import "testing"

func Test_FindGCD1(t *testing.T) {

	nums := []int{2, 5, 6, 9, 10}
	res := FindGCD(nums)
	t.Logf("Expected: %s, got: %s", "ABC", res)
}

func Test_FindGCD2(t *testing.T) {
	nums := []int{7, 5, 6, 8, 3}
	res := FindGCD(nums)
	t.Logf("Expected: %s, got: %s", "AB", res)
}

func Test_FindGCD3(t *testing.T) {
	nums := []int{3, 3}
	res := FindGCD(nums)
	t.Logf("Expected: %s, got: %s", "", res)
}

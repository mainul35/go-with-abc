package leetcode

import "testing"

func Test_abcabcbb(t *testing.T) {
	var s string = "abcabcbb"
	out := LengthOfLongestSubstring(s)
	t.Logf("Expected: %d, Actual: %d", 3, out)
}

func Test_bbbbb(t *testing.T) {
	var s string = "bbbbb"
	out := LengthOfLongestSubstring(s)
	t.Logf("Expected: %d, Actual: %d", 1, out)
}

func Test_pwwkew(t *testing.T) {
	var s string = "pwwkew"
	out := LengthOfLongestSubstring(s)
	t.Logf("Expected: %d, Actual: %d", 3, out)
}

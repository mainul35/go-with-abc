package leetcode

import "testing"

func TestReverseWords1(t *testing.T) {
	s := "the sky is blue"
	res := ReverseWords(s)
	t.Logf("Expected: %s, Actual: %s", "blue is sky the", res)
}

func TestReverseWords2(t *testing.T) {
	s := "  hello world  "
	res := ReverseWords(s)
	t.Logf("Expected: %s, Actual: %s", "world hello", res)
}

func TestReverseWords3(t *testing.T) {
	s := "a good   example"
	res := ReverseWords(s)
	t.Logf("Expected: %s, Actual: %s", "example good a", res)
}

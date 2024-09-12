package leetcode

import "testing"

func Test_ReverseVowelsOfAString1(t *testing.T) {
	str := "hello"
	res := ReverseVowels(str)
	t.Logf("Expected: %s, Actual: %s", "holle", res)
}

func Test_ReverseVowelsOfAString2(t *testing.T) {
	str := "leetcode"
	res := ReverseVowels(str)
	t.Logf("Expected: %s, Actual: %s", "leotcede", res)
}

package leetcode

import "testing"

func Test_gcdString1(t *testing.T) {

	str1 := "ABCABC"
	str2 := "ABC"
	res := GcdOfStrings(str1, str2)
	t.Logf("Expected: %s, got: %s", "ABC", res)
}

func Test_gcdString2(t *testing.T) {

	str1 := "ABABAB"
	str2 := "ABAB"
	res := GcdOfStrings(str1, str2)
	t.Logf("Expected: %s, got: %s", "AB", res)
}

func Test_gcdString3(t *testing.T) {

	str1 := "LEET"
	str2 := "CODE"
	res := GcdOfStrings(str1, str2)
	t.Logf("Expected: %s, got: %s", "", res)
}

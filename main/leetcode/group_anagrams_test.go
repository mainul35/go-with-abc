package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams_withValues(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	expected := [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}
	t.Logf("Expected %v, got %v", expected, res)
}

func Test_groupAnagrams_withZeroLengthList(t *testing.T) {
	strs := []string{""}
	res := groupAnagrams(strs)
	t.Logf("Result: %v", res)

	expected := [][]string{{""}}
	// Compare the expected and result slices
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, but got %v", expected, res)
	}
}

func Test_groupAnagrams_withOneLengthList(t *testing.T) {
	strs := []string{"a"}
	res := groupAnagrams(strs)
	t.Logf("Result: %v", res)

	expected := [][]string{{"a"}}
	// Compare the expected and result slices
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v, but got %v", expected, res)
	}
}

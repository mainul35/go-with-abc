package leetcode

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	chars1 := []rune(s)
	chars2 := []rune(t)
	sort.Slice(chars1, func(i, j int) bool {
		return chars1[i] < chars1[j]
	})

	sort.Slice(chars2, func(i, j int) bool {
		return chars2[i] < chars2[j]
	})

	s1 := string(chars1)
	s2 := string(chars2)

	return s1 == s2
}

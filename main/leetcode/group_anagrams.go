package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	mapStrs := make(map[string][]string)

	for _, word := range strs {
		sortedStr := sortString(word)
		if _, found := mapStrs[sortedStr]; !found {
			mapStrs[sortedStr] = []string{}
		}
		mapStrs[sortedStr] = append(mapStrs[sortedStr], word)
	}

	result := [][]string{}
	for _, v := range mapStrs {
		result = append(result, v)
	}

	return result
}

func sortString(inputString string) string {
	chars := strings.Split(inputString, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

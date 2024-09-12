package leetcode

import (
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")

	var i int = 0
	j := len(words) - 1

	for i < j {
		if !isEmpty(words[i]) && !isEmpty(words[j]) {
			temp := words[i]
			words[i] = words[j]
			words[j] = temp
			i++
			j--
		} else if isEmpty(words[i]) {
			i++
		} else {
			j--
		}
	}

	s = ""
	for _, word := range words {
		if word != "" {
			s += word + " "
		}
	}
	return strings.Trim(s, " ")
}

func isEmpty(s string) bool {
	return s == ""
}

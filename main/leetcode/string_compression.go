package leetcode

import (
	"strconv"
)

func Compress(chars []byte) int {
	i, j := 0, 0
	count := 1

	for {
		if i+1 < len(chars) && chars[i] == chars[i+1] {
			i++
			count++
		} else {
			chars[j] = chars[i]
			j++
			if count > 1 {
				countStr := strconv.Itoa(count)
				for k := 0; k < len(countStr); k++ {
					chars[j] = countStr[k]
					j++
				}
			}
			count = 1
			if i+1 == len(chars) {
				break
			}
			i++
		}
	}
	return j
}

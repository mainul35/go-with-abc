package leetcode

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	//nums := []int{3,2,4}
	//target := 6
	fmt.Println(twoSum(nums, target))
}

// key = item, value = index
// nums := []int{3,2,4}
// target := 6
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, item := range nums {
		if value, ok := m[target-item]; ok {
			return []int{i, value}
		}
		m[item] = i
	}
	return []int{}
}

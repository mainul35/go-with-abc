package leetcode

import "testing"

func Test_input_for_number_of_islands_1(t *testing.T) {
	var arr = [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	var res = numIslands(arr)
	t.Logf("Expected: %d, Actual: %d", 1, res)
}

func Test_input_for_number_of_islands_2(t *testing.T) {
	var arr = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	var res = numIslands(arr)
	t.Logf("Expected: %d, Actual: %d", 3, res)
}

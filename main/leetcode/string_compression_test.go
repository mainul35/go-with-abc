package leetcode

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
		{[]byte{'a', 'a'}, 2},
		{[]byte{'a', 'b'}, 2},
		{[]byte{'a', 'a', 'b'}, 3},
	}

	for _, test := range tests {
		result := Compress(test.input)
		if result != test.expected {
			t.Errorf("compress(%v) = %d; expected %d", test.input, result, test.expected)
		} else {
			fmt.Printf("compress(%v) = %d\n", test.input, result)
		}
	}
}

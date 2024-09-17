package leetcode

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"listen", "silent", true},
		{"hello", "world", false},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abc", "cba", true},
		{"abc", "abcd", false},
	}

	for _, test := range tests {
		result := IsAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("isAnagram(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}

package leetcode

import "fmt"

// Define a set using a map with empty struct values
type Set map[rune]struct{}

// Add adds an element to the set
func (s Set) Add(value rune) {
	s[value] = struct{}{}
}

// Remove removes an element from the set
func (s Set) Remove(value rune) {
	delete(s, value)
}

// Contains checks if the set contains the given element
func (s Set) Contains(value rune) bool {
	_, exists := s[value]
	return exists
}

// Size returns the number of elements in the set
func (s Set) Size() int {
	return len(s)
}

// Print prints all elements in the set
func (s Set) Print() {
	for key := range s {
		fmt.Println(key)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// lengthOfLongestSubstring returns the length of the longest substring without repeating characters
func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	charSet := make(Set)
	left := 0

	for right, char := range s {
		if !charSet.Contains(char) {
			charSet.Add(char)
			maxLen = max(maxLen, right-left+1)
		} else {
			for charSet.Contains(char) {
				charSet.Remove(rune(s[left]))
				left++
			}
			charSet.Add(char)
		}
	}
	return maxLen
}

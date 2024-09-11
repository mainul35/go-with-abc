package leetcode

import "testing"

func Test_KidsWithCandies1(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	expectedOutput := []bool{true, true, true, false, true}
	output := KidsWithCandies(candies, extraCandies)
	passed := true
	// Test the values of the array
	for i, v := range output {
		if v != expectedOutput[i] {
			t.Logf("At index %d, expected: %t, but got: %t", i, expectedOutput[i], v)
			passed = false
			t.Fail()
		}
	}
	if passed != true {
		t.Logf("Test Failed")
	} else {
		t.Logf("Test Passed")
	}
}

func Test_KidsWithCandies2(t *testing.T) {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	expectedOutput := []bool{true, false, false, false, false}
	output := KidsWithCandies(candies, extraCandies)
	passed := true
	// Test the values of the array
	for i, v := range output {
		if v != expectedOutput[i] {
			t.Logf("At index %d, expected: %t, but got: %t", i, expectedOutput[i], v)
			passed = false
			t.Fail()
		}
	}
	if passed != true {
		t.Logf("Test Failed")
	} else {
		t.Logf("Test Passed")
	}
}

func Test_KidsWithCandies3(t *testing.T) {
	candies := []int{12, 1, 12}
	extraCandies := 10
	expectedOutput := []bool{true, false, true}
	output := KidsWithCandies(candies, extraCandies)
	passed := true
	// Test the values of the array
	for i, v := range output {
		if v != expectedOutput[i] {
			t.Logf("At index %d, expected: %t, but got: %t", i, expectedOutput[i], v)
			passed = false
			t.Fail()
		}
	}
	if passed != true {
		t.Logf("Test Failed")
	} else {
		t.Logf("Test Passed")
	}
}

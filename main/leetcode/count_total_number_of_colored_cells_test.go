package leetcode

import "testing"

func Test_coloredCells1(t *testing.T) {
	res := ColoredCells(1)
	t.Logf("Expected: %d, got: %d", 1, res)
}

func Test_coloredCells2(t *testing.T) {
	res := ColoredCells(2)
	t.Logf("Expected: %d, got: %d", 5, res)
}

func Test_coloredCells3(t *testing.T) {
	res := ColoredCells(3)
	t.Logf("Expected: %d, got: %d", 13, res)
}

func Test_coloredCells4(t *testing.T) {
	res := ColoredCells(4)
	t.Logf("Expected: %d, got: %d", 25, res)
}

func Test_coloredCells5(t *testing.T) {
	res := ColoredCells(5)
	t.Logf("Expected: %d, got: %d", 41, res)
}

func Test_coloredCells6(t *testing.T) {
	res := ColoredCells(6)
	t.Logf("Expected: %d, got: %d", 61, res)
}

func Test_coloredCells7(t *testing.T) {
	res := ColoredCells(7)
	t.Logf("Expected: %d, got: %d", 85, res)
}

package leetcode

import "testing"

func TestCanPlaceFlowers1(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	res := CanPlaceFlowers(flowerbed, n)
	t.Logf("%v", res)
}

func TestCanPlaceFlowers2(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	res := CanPlaceFlowers(flowerbed, n)
	t.Logf("%v", res)
}

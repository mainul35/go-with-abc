package leetcode

func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 1
	for _, qty := range candies {
		if max < qty {
			max = qty
		}
	}
	n := len(candies)
	boolArr := make([]bool, n)
	for i, qty := range candies {
		if qty+extraCandies >= max {
			boolArr[i] = true
		} else {
			boolArr[i] = false
		}
	}

	return boolArr
}

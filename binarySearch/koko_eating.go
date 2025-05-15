package binarysearch

import "math"

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	possibleToEat := func(speed int) bool {
		var hours int
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(speed)))
		}
		return hours <= h
	}

	var result int
	for left <= right {
		mid := left + ((right - left) / 2)
		if possibleToEat(mid) {
			right = mid - 1
			result = mid
		} else {
			left = mid + 1
		}
	}

	return result
}

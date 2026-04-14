package binarysearch

import (
	"math"
	"sort"
)

// https://neetcode.io/problems/eating-bananas/question
// the idea we test the possible speeds from the minimum to the maximum value in the piles
// we test the mid value and if it's possible to eat, we save the mid value as the result
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	minSpeed := 1
	maxSpeed := piles[len(piles)-1]

	isPossibleToEat := func(speed int) bool {
		var hours int
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(speed)))
		}

		return hours <= h
	}

	var result int
	for minSpeed <= maxSpeed {
		currMidSpeed := minSpeed + ((maxSpeed - minSpeed) / 2)
		if isPossibleToEat(currMidSpeed) {
			// even tho the current value is possible, we still need to check the minimum value
			maxSpeed = currMidSpeed - 1
			// but we still save the current value as possible result
			result = currMidSpeed
		} else {
			minSpeed = currMidSpeed + 1
		}
	}

	return result
}

func minEatingSpeed2(piles []int, h int) int {
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

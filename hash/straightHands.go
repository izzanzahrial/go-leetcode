package hash

import "sort"

// https://leetcode.com/problems/hand-of-straights/
// still looks ugly

func isNStraightHand(hand []int, groupSize int) bool {
	hashMap := make(map[int]int)

	for _, val := range hand {
		hashMap[val]++
	}

	sort.Ints(hand)
	for _, val := range hand {

		if hashMap[val] > 0 {
			counter := groupSize
			valCount := val

			for counter > 0 {
				if hashMap[valCount] > 0 {
					hashMap[valCount] -= 1
					counter -= 1
				} else {
					return false
				}
				valCount++
			}
		}
	}

	return true
}

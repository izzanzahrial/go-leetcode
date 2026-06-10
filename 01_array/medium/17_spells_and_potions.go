package array

import "sort"

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/?envType=study-plan-v2&envId=leetcode-75
// the idea is to sort the potions array
// then check by searching the mid of the potions array
// by doing that we already eliminate alot of the potions
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	res := make([]int, len(spells))

	for i, s := range spells {
		s64 := int64(s)
		need := (success + s64 - 1) / s64

		left, right := 0, m
		for left < right {
			mid := (left + right) / 2
			if int64(potions[mid]) < need {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// we got the left index of minimum potions strength needed
		// so just subtract the left index from the total length of potions
		res[i] = m - left
	}
	return res
}

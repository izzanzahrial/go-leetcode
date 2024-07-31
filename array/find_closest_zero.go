package array

import "math"

// https://leetcode.com/problems/find-closest-number-to-zero/
func findClosestNumber(nums []int) int {
	var closest int
	curr := math.MaxInt

	for _, num := range nums {
		absCurr := int(math.Abs(float64(num)))
		if absCurr <= curr {
			if closest > num && absCurr == curr {
				continue
			}
			curr = absCurr
			closest = num
		}
	}

	return closest
}

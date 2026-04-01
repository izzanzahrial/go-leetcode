package array

// https://leetcode.com/problems/max-consecutive-ones/description/
func FindMaxConsecutiveOnes(nums []int) int {
	var max, curr int
	for _, num := range nums {
		if num == 1 {
			curr++
		} else {
			curr = 0
		}

		if curr > max {
			max = curr
		}
	}

	return max
}

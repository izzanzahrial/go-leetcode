package array

// Can make it better using sliding window
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/?envType=study-plan-v2&envId=leetcode-75
func LongestSubarray(nums []int) int {
	var max int
	var isZero bool

	for idx, val := range nums {
		var curr, deleted int
		if val == 0 {
			isZero = true
		}

		if val == 1 {
			for i := idx; i < len(nums); i++ {
				if nums[i] == 1 {
					curr += 1
				}

				if nums[i] == 0 && deleted == 0 {
					deleted = 1
					continue
				}

				if nums[i] == 0 && deleted > 0 {
					break
				}
			}
		}

		if curr > max {
			max = curr

			if deleted == 0 && !isZero {
				max -= 1
			}
		}
	}

	return max
}

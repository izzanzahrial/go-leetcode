package pointer

import "sort"

// https://leetcode.com/problems/4sum/description/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	// -3 stops three elements before the end because we need at least three more numbers
	for i := 0; i < len(nums)-3; i++ {
		// i > 0: This ensures we don't check nums[i-1] when i is 0 (which would be an out-of-bounds access)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// -2 stops two elements before the end because we need at least two more numbers
		for j := i + 1; j < len(nums)-2; j++ {
			// The check j > i+1 ensures that we are only skipping nums[j] if it's a duplicate of the previous value
			// considered for the j-th position, not a duplicate of nums[i]
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1
			for left < right {
				currSum := nums[i] + nums[j] + nums[left] + nums[right]
				if currSum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					left++
					right--
					// Move left pointer forward as long as it points to a duplicate of the number just used for the third element
					// This ensures that if nums[left] was X, and the next element is also X, we don't form another quadruplet
					// with the same nums[i], nums[j], and this new X
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					// Similarly, move right pointer backward to skip duplicates for the fourth element
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if currSum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}

package array

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func SearchRange(nums []int, target int) []int {
	var result []int
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + ((right - left) / 2)
		val := nums[mid]

		if val == target {
			// check left side
			for i := mid - 1; i >= 0; i-- {
				if nums[i] != target {
					result = append(result, i+1)
					break
				}
			}

			// check right side
			for j := mid + 1; j < len(nums); j++ {
				if nums[j] != target {
					result = append(result, j-1)
					break
				}
			}

			return result
		}

		if val > target {
			right = mid - 1
		}

		if val < target {
			left = mid + 1
		}
	}

	return []int{-1, -1}
}

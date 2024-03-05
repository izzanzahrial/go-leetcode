package array

// https://leetcode.com/problems/search-insert-position/description/
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + ((right - left) / 2)
		val := nums[mid]

		if val == target {
			return mid
		}

		if val < target {
			left = mid + 1
		}

		if val > target {
			right = mid - 1
		}
	}

	return left
}

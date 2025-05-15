package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		// if the nums[mid] less than nums[right], meaning the number is smaller
		// so the minimum number, might be in the left side
		// change the right into mid
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

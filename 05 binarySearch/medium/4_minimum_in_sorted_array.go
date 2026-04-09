package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		// check if is the array is normal
		// meaning the right side of mid is still bigger than the mid itself
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

package binarysearch

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if target > nums[middle] || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}

package binarysearch

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		currNum := nums[mid]
		if currNum == target {
			return mid
		}

		// if the middle number is smaller or equal to right number
		if currNum <= nums[right] {
			// if the right number bigger or equal to the target
			// and target bigger than middle number
			// meaning the target should be bigger than middle number
			// so the left index should be move to the middle + 1 index
			// the target should be between right and middle number
			if nums[right] >= target && target > currNum {
				left = mid + 1
			} else {
				// else, meaning the target should be smaller than middle number
				// so the target should be between middle and left number
				right = mid - 1
			}
		} else {
			// if the left number smaller or equal to the target
			// and the target smaller than middle number
			// meaning the target should be smaller than middle number
			// so the right index should be move to the middle - 1 index
			// the target should be between middle and left number
			if nums[left] <= target && target < currNum {
				right = mid - 1
			} else {
				// else, meaning the target should be bigger than middle number
				// so the target should be between middle and right number
				left = mid + 1
			}
		}
	}

	return -1
}

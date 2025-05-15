package binarysearch

// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func search3(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// The key is here: Handling duplicates that make it hard to determine the sorted half
		if nums[left] == nums[mid] {
			// If nums[left] and nums[mid] are the same, we can't tell if the
			// segment from nums[left] to nums[mid] is sorted or if the pivot
			// lies within it.
			// For example:
			// [1, 0, 1, 1, 1] -> left=0, mid=2. nums[left]=1, nums[mid]=1. Left part [1,0,1] is not sorted.
			// [1, 1, 1, 0, 1] -> left=0, mid=2. nums[left]=1, nums[mid]=1. Left part [1,1,1] IS sorted.
			// Since nums[mid] wasn't the target, and nums[left] is the same as nums[mid],
			// nums[left] also cannot be the target (if it were, mid would have been the target).
			// So, we can safely advance 'left' by one to try and break the tie and get more information
			// in the next iteration. This essentially "skips" the duplicate nums[left].
			left++
			continue // Restart the loop with the new 'left'
		}

		// If in they key section where you modified if the nums is the same and using the left index
		// you have to use the left index again as conditional
		// If you switch the main decision to nums[mid] <= nums[right], you are now relying on nums[mid]
		// and nums[right] to determine the sorted half. The left++ operation doesn't directly "clean up" potential ambiguities
		// between mid and right in the same way it does for left and mid
		// For example, if nums[left] != nums[mid] but nums[mid] == nums[right]:
		// nums = [5, 1, 2, 2, 2], target = 5
		// left=0(5), right=4(2), mid=2(2).
		// nums[mid] != target.
		// nums[left] == nums[mid] (5 == 2) is false. No left++.
		// Your condition: nums[mid] <= nums[right] (2 <= 2) is TRUE. So you assume right half [2,2,2] is sorted.
		// Target 5 in [2,2,2]? No (target > nums[mid] && target <= nums[right] -> 5 > 2 && 5 <= 2 is false).
		// So you'd set right = mid - 1 (i.e., right = 1). Search space becomes [5,1]. This is correct.
		// Original condition: nums[left] <= nums[mid] (5 <= 2) is FALSE. So it assumes right half [2,2,2] is sorted.
		// Target 5 in [2,2,2]? No (target > nums[mid] && target <= nums[right] -> 5 > 2 && 5 <= 2 is false).
		// So it'd set right = mid - 1. This is also correct.
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

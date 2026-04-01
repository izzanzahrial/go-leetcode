package array

// https://leetcode.com/problems/semi-ordered-permutation/
func FirstAndLastValueIdx(nums []int) int {
	var firstIdx, lastIdx int
	for left, right := 0, len(nums)-1; left <= right; left, right = left+1, right-1 {
		if nums[left] == 1 {
			firstIdx = left
		}

		if nums[left] == len(nums) {
			lastIdx = left
		}

		if nums[right] == 1 {
			firstIdx = right
		}

		if nums[right] == len(nums) {
			lastIdx = right
		}
	}

	var result int
	// count the step for the firstIdx to be in the actual 0 idx (first idx)
	result = firstIdx
	// count the step for the lastIdx to be in the actual last idx (len(nums)-1)
	result += len(nums) - 1 - lastIdx
	// if the firstIdx and the lastIdx cross over, meaning first at the right side
	// then lastIdx at the left side, we only need 1 operations to swap them
	// so need to recude the result by 1
	if firstIdx > lastIdx {
		result--
	}

	return result
}

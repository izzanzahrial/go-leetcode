package array

// https://leetcode.com/problems/find-peak-element/description/
func FindPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if isPeak(i, nums) {
			return i
		}
	}

	return 0
}

func isPeak(idx int, nums []int) bool {
	if idx == 0 {
		return nums[idx] > nums[idx+1]
	}

	if idx == len(nums)-1 {
		return nums[idx] > nums[idx-1]
	}

	if nums[idx] > nums[idx+1] && nums[idx] > nums[idx-1] {
		return true
	}

	return false
}

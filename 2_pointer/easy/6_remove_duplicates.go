package pointer

// https://neetcode.io/problems/remove-duplicates-from-sorted-array/question
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	pointer1, pointer2 := 0, 1
	for ; pointer2 < len(nums); pointer2++ {
		if nums[pointer1] == nums[pointer2] {
			nums = append(nums[:pointer2], nums[pointer2+1:]...)
			pointer2--
		} else {
			pointer1++
		}
	}

	return len(nums)
}

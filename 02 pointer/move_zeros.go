package pointer

func MoveZeroes(nums []int) {
	pivot := 0
	search := 0
	for pivot < len(nums) {
		if nums[pivot] != 0 {
			nums[search], nums[pivot] = nums[pivot], nums[search]
		}

		if nums[search] != 0 {
			search += 1
		}

		pivot += 1
	}
}

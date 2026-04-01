package array

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	if right <= left+1 {
		if right == left+1 && nums[right] < nums[left] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		return
	}

	j := partition(nums, left, right)
	quickSort(nums, left, j-1)
	quickSort(nums, j+1, right)
}

func partition(nums []int, left int, right int) int {
	mid := left + (right-left)/2
	nums[mid], nums[left+1] = nums[left+1], nums[mid]

	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	}

	if nums[left+1] > nums[right] {
		nums[left+1], nums[right] = nums[right], nums[left+1]
	}

	if nums[left] > nums[left+1] {
		nums[left], nums[left+1] = nums[left+1], nums[left]
	}

	pivot := nums[left+1]
	i := left + 1
	j := right

	for {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i > j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[left+1], nums[j] = nums[j], nums[left+1]
	return j
}

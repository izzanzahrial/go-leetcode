package backtrack

// https://leetcode.com/problems/permutations
func permute(nums []int) [][]int {
	var result [][]int

	var recursive func(idx int, nums []int)
	recursive = func(idx int, arr []int) {
		if idx == len(nums) {
			// create a deep copy, because if it's not
			// then later on the curr get modified
			// then the curr within the result also get modified
			temp := append([]int{}, nums...)
			result = append(result, temp)
			return
		}

		for i := idx; i < len(nums); i++ {
			nums[i], nums[idx] = nums[idx], nums[i]
			recursive(idx+1, nums)
			// bactrack
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	recursive(0, nums)
	return result
}

// version 2
func permute2(nums []int) [][]int {
	var result [][]int

	recursive([]int{}, nums, &result)

	return result
}

func recursive(curr, nums []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, curr)
	}

	for i, val := range nums {
		curr1 := nums[:i]
		curr2 := nums[i+1:]
		newCurr := append(curr, val)
		newNums := append(append([]int{}, curr1...), curr2...)
		recursive(newCurr, newNums, result)
	}
}

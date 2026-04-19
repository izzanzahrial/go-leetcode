package backtrack

// https://leetcode.com/problems/subsets/description/
func subsets(nums []int) [][]int {
	var result [][]int

	var recursive func(curr []int, idx int)
	recursive = func(curr []int, idx int) {
		// if we are using this, and in the later it get update
		// the original curr also get updated, because it is not a deep copy
		// result = append(result, curr)

		// create a copy
		temp := make([]int, len(curr))
		copy(temp, curr)
		result = append(result, temp)

		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			recursive(curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	recursive([]int{}, 0)
	return result
}

package backtrack

import "sort"

// https://leetcode.com/problems/subsets
func subsetsWithDup(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)
	var recursive func(idx int, curr []int)
	recursive = func(idx int, curr []int) {
		// create a deep copy, because if it's not
		// then later on the curr get modified
		// then the curr within the result also get modified
		temp := make([]int, len(curr))
		copy(temp, curr)
		result = append(result, temp)

		for i := idx; i < len(nums); i++ {
			// to make sure we dont get duplicate
			// it's basically checking within the loop within the same recursive function
			// not the one that going more in depth
			// example the incomming idx is = 0 and i = idx, in the next loop i will become 1
			// but not the 'i' that going into recursive(i+1)
			// if the current candidates is the same as the prev one then we skip it
			if i > idx && nums[i] == nums[i-1] {
				continue
			}

			curr = append(curr, nums[i])
			recursive(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	recursive(0, []int{})
	return result
}

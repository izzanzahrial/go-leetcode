package backtrack

import "sort"

// https://leetcode.com/problems/combination-sum-ii/
// https://www.youtube.com/watch?v=rSA3t6BDDwg&ab_channel=NeetCode
func CombinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(0, 0, target, candidates, []int{}, &result)
	return result
}

func backtrack(idx, total, target int, candidates, curr []int, result *[][]int) {
	if total == target {
		valid := make([]int, len(curr))
		copy(valid, curr)
		*result = append(*result, valid)
		return
	}

	if total > target {
		return
	}

	var prev int
	for i := idx; i < len(candidates); i++ {
		// prevents using the same value multiple times at the same decision level
		if candidates[i] == prev {
			continue
		}

		curr = append(curr, candidates[i])
		backtrack(i+1, total+candidates[i], target, candidates, curr, result)
		curr = curr[:len(curr)-1]
		prev = candidates[i]
	}
}

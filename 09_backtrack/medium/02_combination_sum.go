package backtrack

import "sort"

// https://leetcode.com/problems/combination-sum/
// https://www.youtube.com/watch?v=GBKI9VSKdGg&ab_channel=NeetCode
func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return nil
	}

	var result [][]int
	var curr []int
	backtrack(0, 0, target, curr, candidates, &result)
	return result
}

func backtrack(idx, currTotal, target int, curr, candidates []int, result *[][]int) {
	// check if the current total > target
	// and check to make sure we don't go the last index + 1
	if currTotal > target || idx >= len(candidates) {
		return
	}

	if currTotal == target {
		// to make sure the current result doesn't get updated by the previous call
		// create copy of the current result
		currCopy := make([]int, len(curr))
		copy(currCopy, curr)
		*result = append(*result, currCopy)
		return
	}

	curr = append(curr, candidates[idx])
	backtrack(idx, currTotal+candidates[idx], target, curr, candidates, result)

	// backtrack then go to the next index
	curr = curr[:len(curr)-1]
	backtrack(idx+1, currTotal, target, curr, candidates, result)
}

// https://www.youtube.com/watch?v=rSA3t6BDDwg&ab_channel=NeetCode
// version 2
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(0, 0, target, candidates, []int{}, &result)
	return result
}

func backtrack2(idx, total, target int, candidates, curr []int, result *[][]int) {
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
		backtrack2(i+1, total+candidates[i], target, candidates, curr, result)
		curr = curr[:len(curr)-1]
		prev = candidates[i]
	}
}

// version 3
func combinationSum3(nums []int, target int) [][]int {
	var result [][]int

	var recursive func(idx, curr int, currArr []int)
	recursive = func(idx, curr int, currArr []int) {
		if curr == target {
			// create a deep copy, because if it's not
			// then later on the currArr get modified
			// then the currArr within the result also get modified
			temp := make([]int, len(currArr))
			copy(temp, currArr)
			result = append(result, temp)
			return
		}

		for i := idx; i < len(nums); i++ {
			if curr+nums[i] > target {
				return
			}
			currArr = append(currArr, nums[i])
			recursive(i, curr+nums[i], currArr)

			// backtrack
			currArr = currArr[:len(currArr)-1]
		}
	}

	recursive(0, 0, []int{})
	return result
}

package backtrack

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

func backtrack1(idx, currTotal, target int, curr, candidates []int, result *[][]int) {
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

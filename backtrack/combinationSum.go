package backtrack

// https://leetcode.com/problems/combination-sum/
func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var curr []int

	backtrack1(0, 0, target, candidates, curr, &result)

	return result
}

func backtrack1(i, total, target int, candidates, curr []int, result *[][]int) {
	if total == target {
		valid := make([]int, len(curr))
		copy(valid, curr)
		*result = append(*result, valid)
		return
	}

	if i >= len(candidates) || total > target {
		return
	}

	curr = append(curr, candidates[i])
	backtrack1(i, total+candidates[i], target, candidates, curr, result)
	curr = curr[:len(curr)-1]
	backtrack1(i+1, total, target, candidates, curr, result)
}

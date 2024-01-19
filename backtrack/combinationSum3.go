package backtrack

// https://leetcode.com/problems/combination-sum-iii/?envType=study-plan-v2&envId=leetcode-75
// idk why this is wrong, even though i debug the result is right, but the print out result still wrong
func CombinationSum3(k int, n int) [][]int {
	var result [][]int

	// i is the current starting number
	var backtrack func(i, k, total int, curr []int)
	backtrack = func(i, k, total int, curr []int) {
		if total > n {
			return
		}

		if total == n && k == 0 {
			result = append(result, curr)
			return
		}

		if k == 0 {
			return
		}

		for j := i; j <= 9; j++ {
			curr = append(curr, j)
			newTotal := total + j
			backtrack(j+1, k-1, newTotal, curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1, k, 0, []int{})

	return result
}

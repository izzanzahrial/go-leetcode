package backtrack

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	var result [][]int

	var recursive func(idx int, curr []int)
	recursive = func(idx int, curr []int) {
		if len(curr) == k {
			// create a deep copy, because if it's not
			// then later on the curr get modified
			// then the curr within the result also get modified
			temp := make([]int, len(curr))
			copy(temp, curr)
			result = append(result, temp)
			return
		}

		for i := idx; i <= n; i++ {
			curr = append(curr, i)

			// if the current length is greater than k
			// then break the loop, because we don't want to go further
			if len(curr) > k {
				curr = curr[:len(curr)-1]
				break
			}

			recursive(i+1, curr)

			// backtrack
			curr = curr[:len(curr)-1]
		}
	}

	recursive(1, []int{})
	return result
}

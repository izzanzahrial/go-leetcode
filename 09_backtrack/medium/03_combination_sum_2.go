package backtrack

import "sort"

// https://leetcode.com/problems/combination-sum-ii
func combinationsSum2(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)
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

		for i := idx; i < len(candidates); i++ {
			// to make sure we dont get duplicate
			// it's basically checking within the loop within the same recursive function
			// not the one that going more in depth
			// example the incomming idx is = 0 and i = idx, in the next loop i will become 1
			// but not the 'i' that going into recursive(i+1)
			// if the current candidates is the same as the prev one then we skip it
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			if curr+candidates[i] > target {
				break
			}

			currArr = append(currArr, candidates[i])
			recursive(i+1, curr+candidates[i], currArr)

			// backtrack
			currArr = currArr[:len(currArr)-1]
		}
	}

	recursive(0, 0, []int{})
	return result
}

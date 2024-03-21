package backtrack

// https://leetcode.com/problems/all-paths-from-source-to-target/
func AllPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	curr := []int{0}
	backtrack4(0, curr, graph, &result)
	return result
}

func backtrack4(num int, curr []int, graph [][]int, result *[][]int) {
	if curr[len(curr)-1] == len(graph)-1 {
		valid := make([]int, len(curr))
		copy(valid, curr)
		*result = append(*result, valid)
		return
	}

	for _, val := range graph[num] {
		curr = append(curr, val)
		backtrack4(val, curr, graph, result)
		curr = curr[:len(curr)-1]
	}
}

package graph

// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
func countComponents(n int, edges [][]int) int {
	edgeMap := make(map[int][]int)
	for _, edge := range edges {
		edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
		edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
	}

	visited := make([]bool, n)

	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true
		for _, nei := range edgeMap[node] {
			dfs(nei)
		}
	}

	var result int
	for i := range n {
		if visited[i] {
			continue
		}

		dfs(i)
		result++
	}

	return result
}

package graph

// https://leetcode.com/problems/graph-valid-tree/description/
func validTree(n int, edges [][]int) bool {
	if len(edges) > n-1 {
		return false
	}

	// make a list of edges that are connected to each others
	adj := make([][]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		adj[i] = append(adj[i], j)
		adj[j] = append(adj[j], i)
	}

	visited := make(map[int]struct{})

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		if _, exists := visited[node]; exists {
			return false
		}

		visited[node] = struct{}{}
		// check all the nodes that are connected to the current node
		for _, nei := range adj[node] {
			if nei == parent {
				continue
			}

			if !dfs(nei, node) {
				return false
			}
		}

		return true
	}

	// check the node 0 and -1 (because 0 doesnt have parent)
	return dfs(0, -1) && len(visited) == n
}

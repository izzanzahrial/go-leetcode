package dfs

import (
	"math"

	"github.com/izzanzahrial/go-leetcode/dfs"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *dfs.TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := math.MinInt
	var countDepth func(currDepth int, root *dfs.TreeNode)
	countDepth = func(currDepth int, root *dfs.TreeNode) {
		if root.Left != nil {
			countDepth(currDepth+1, root.Left)
		}

		if root.Right != nil {
			countDepth(currDepth+1, root.Right)
		}

		maxDepth = max(maxDepth, currDepth)
	}
	countDepth(1, root)

	return maxDepth
}

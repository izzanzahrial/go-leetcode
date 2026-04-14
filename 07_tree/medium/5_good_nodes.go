package tree

import (
	"math"

	tree "github.com/izzanzahrial/go-leetcode/07_tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://neetcode.io/problems/count-good-nodes-in-binary-tree/question
func goodNodes(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, math.MinInt)
}

func dfs(root *tree.TreeNode, currMax int) int {
	if root == nil {
		return 0
	}

	var result int
	if root.Val >= currMax {
		currMax = max(currMax, root.Val)
		result++
	}

	left := dfs(root.Left, currMax)
	right := dfs(root.Right, currMax)

	return result + left + right
}

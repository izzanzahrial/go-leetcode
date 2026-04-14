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
// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *tree.TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(root *tree.TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Val <= left || root.Val >= right {
		return false
	}

	return valid(root.Left, left, root.Val) && valid(root.Right, root.Val, right)
}

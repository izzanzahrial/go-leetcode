package tree

import "github.com/izzanzahrial/go-leetcode/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/balanced-binary-tree/description/
func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.Left)
	right := height(root.Right)
	if abs(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

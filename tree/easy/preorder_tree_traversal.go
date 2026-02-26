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
// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	var preorder func(root *tree.TreeNode)
	preorder = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)

	return result
}

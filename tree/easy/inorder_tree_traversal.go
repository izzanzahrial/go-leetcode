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
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
func inorderTraversal(root *tree.TreeNode) []int {
	var result []int

	helper(root, &result)

	return result
}

func helper(root *tree.TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	helper(root.Left, arr)
	*arr = append(*arr, root.Val)
	helper(root.Right, arr)
}

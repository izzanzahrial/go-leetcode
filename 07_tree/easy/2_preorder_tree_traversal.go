package tree

import tree "github.com/izzanzahrial/go-leetcode/07_tree"

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

	result := make([]int, 0)

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

func preorderTraversal2(root *tree.TreeNode) []int {
	result := make([]int, 0)

	traverse(root, &result)

	return result
}

func traverse(root *tree.TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)
	traverse(root.Left, arr)
	traverse(root.Right, arr)
}

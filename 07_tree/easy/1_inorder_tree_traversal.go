package tree

import (
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
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
func inorderTraversal(root *tree.TreeNode) []int {
	result := make([]int, 0)

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

// func inorderTraversal(root *tree.TreeNode) []int {
// 	var result []int

// 	var inorder func(root *tree.TreeNode)
// 	inorder = func(root *tree.TreeNode) {
// 		if root == nil {
// 			return
// 		}

// 		inorder(root.Left)
// 		result = append(result, root.Val)
// 		inorder(root.Right)
// 	}

// 	inorder(root)

// 	return result
// }

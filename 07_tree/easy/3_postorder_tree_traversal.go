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
// https://leetcode.com/problems/binary-tree-postorder-traversal/description/
func postorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	var postorder func(root *tree.TreeNode)
	postorder = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		result = append(result, root.Val)
	}
	postorder(root)

	return result
}

func postorderTraversal2(root *tree.TreeNode) []int {
	result := make([]int, 0)

	traverse2(root, &result)

	return result
}

func traverse2(root *tree.TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	traverse2(root.Left, arr)
	traverse2(root.Right, arr)
	*arr = append(*arr, root.Val)
}

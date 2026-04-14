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

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *tree.TreeNode, k int) int {
	var resultArr []int

	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		resultArr = append(resultArr, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return resultArr[k-1]
}

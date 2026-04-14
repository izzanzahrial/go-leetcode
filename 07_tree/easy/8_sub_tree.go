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
// https://leetcode.com/problems/subtree-of-another-tree
func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if checkTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkTree(p, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return checkTree(p.Left, q.Left) && checkTree(p.Right, q.Right)
	}

	return false
}

package linkedlist

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return getCommonAncestor(root, p, q)
}

func getCommonAncestor(node, p, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == p.Val || node.Val == q.Val {
		return node
	}

	leftSubTree := getCommonAncestor(node.Left, p, q)
	rightSubTree := getCommonAncestor(node.Right, p, q)

	if leftSubTree == nil {
		return rightSubTree
	} else if rightSubTree == nil {
		return leftSubTree
	}

	return node
}

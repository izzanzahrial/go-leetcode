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
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75
func lowestCommonAncestor(root *tree.TreeNode, p, q *tree.TreeNode) *tree.TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestor2(root, p, q *tree.TreeNode) *tree.TreeNode {
	return getCommonAncestor(root, p, q)
}

func getCommonAncestor(node, p, q *tree.TreeNode) *tree.TreeNode {
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

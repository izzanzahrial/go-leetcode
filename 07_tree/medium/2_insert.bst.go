package tree

import tree "github.com/izzanzahrial/go-leetcode/07_tree"

/**
 * Definition for a binary tree node.
 * type tree.TreeNode struct {
 *     Val int
 *     Left *tree.TreeNode
 *     Right *tree.TreeNode
 * }
 */
// https://leetcode.com/problems/insert-into-a-binary-search-tree/
func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}

	if val > root.Val && root.Right != nil {
		insertIntoBST(root.Right, val)
	} else if val < root.Val && root.Left != nil {
		insertIntoBST(root.Left, val)
	} else if val > root.Val && root.Right == nil {
		root.Right = &tree.TreeNode{
			Val: val,
		}
	} else if val < root.Val && root.Left == nil {
		root.Left = &tree.TreeNode{
			Val: val,
		}
	}

	return root
}

func insertIntoBST2(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

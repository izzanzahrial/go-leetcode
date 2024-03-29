package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = MergeTrees(root1.Left, root2.Left)
	root.Right = MergeTrees(root1.Right, root2.Right)

	return root
}

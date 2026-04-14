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
func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		depth++
	}

	return depth
}

func maxDepth2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(countDepth(root.Left, 1), countDepth(root.Right, 1))
}

func countDepth(root *tree.TreeNode, curr int) int {
	if root == nil {
		return curr
	}

	curr++
	return max(countDepth(root.Left, curr), countDepth(root.Right, curr))
}

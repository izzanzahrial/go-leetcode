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
// bfs
// https://leetcode.com/problems/binary-tree-right-side-view
// the idea is to use a queue to traverse the tree
// and keep updating the value within the current breath
func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*tree.TreeNode{root}

	for len(queue) > 0 {

		size := len(queue)
		var currNum int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// keep updating the value until the most right node
			currNum = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currNum)
	}

	return result
}

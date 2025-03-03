package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/minimum-depth-of-binary-tree/description/?envType=problem-list-v2&envId=breadth-first-search&
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var minDepth int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			// early return if left & right nodes is nil
			// meaning it is a leaf node
			if curr.Left == nil && curr.Right == nil {
				return minDepth + 1
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		minDepth++
	}

	return minDepth
}

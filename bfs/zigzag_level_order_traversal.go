package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/?envType=problem-list-v2&envId=breadth-first-search&
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	rightFirst := true

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)

		// check if right first from the entire current level
		// using the value from the current queue
		if rightFirst {
			for _, node := range queue {
				level = append(level, node.Val)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				level = append(level, queue[i].Val)
			}
		}

		result = append(result, level)

		// add the next result of current queue
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

		rightFirst = !rightFirst
	}

	return result
}

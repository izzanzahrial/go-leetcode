package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/?envType=problem-list-v2&envId=breadth-first-search&
func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// we need to dereference the len queue by initializing into size
		// because if not then using it in inner for loop
		// the len(queue) gonna keep growing
		size := len(queue)
		var currLevel []int

		// loop for the lenght of current level tree
		// if we use range, we gonna keep looping until the end of tree
		// because we gonna keep adding to the queue
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			currLevel = append(currLevel, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		result = append(result, currLevel)
	}

	return result
}

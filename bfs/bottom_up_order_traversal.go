package bfs

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/?envType=problem-list-v2&envId=breadth-first-search&
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Prepend to avoid reversing
		// add the current level, then add the previous result
		// creating new result only by reversing the level, not the actual value
		result = append([][]int{level}, result...)
	}

	return result
}

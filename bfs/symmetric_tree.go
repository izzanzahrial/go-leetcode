package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/symmetric-tree/description/?envType=problem-list-v2&envId=breadth-first-search&
// bfs
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		// pop two val from the queue
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		} else if (left == nil && right != nil) || (left != nil && right == nil) {
			// check if one of them nil, return false
			return false
		} else if left.Val != right.Val {
			return false
		}

		// append the next symmetric value
		// left node of the current left with right node of the current right
		queue = append(queue, left.Left, right.Right)

		// // right node of the current left with left node of the current right
		queue = append(queue, left.Right, right.Left)
	}

	return true
}

// recursive
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

package bfs

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/?envType=problem-list-v2&envId=breadth-first-search&
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			// validation to make sure not the last node
			// of the current level of tree
			if size-i > 1 {
				curr.Next = queue[0]
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return root
}

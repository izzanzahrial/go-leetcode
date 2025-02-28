package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/same-tree/submissions/1557734428/
// bfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}

	for len(queueP) > 0 && len(queueQ) > 0 {
		nodeP := queueP[0]
		nodeQ := queueQ[0]

		// pop the first element from the queue
		queueP = queueP[1:]
		queueQ = queueQ[1:]

		if nodeP == nil && nodeQ == nil {
			continue
		}

		if nodeP == nil || nodeQ == nil {
			return false
		}

		if nodeP.Val != nodeQ.Val {
			return false
		}

		queueP = append(queueP, nodeP.Left, nodeP.Right)
		queueQ = append(queueQ, nodeQ.Left, nodeQ.Right)
	}

	return len(queueP) == 0 && len(queueQ) == 0
}

// recursive
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right
}

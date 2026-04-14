package tree

import tree "github.com/izzanzahrial/go-leetcode/07_tree"

// https://leetcode.com/problems/same-tree/submissions/1557734428/
// bfs
func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	queueP := []*tree.TreeNode{p}
	queueQ := []*tree.TreeNode{q}

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
func isSameTree2(p *tree.TreeNode, q *tree.TreeNode) bool {
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

// recursive
func isSameTree3(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}

// dfs
func isSameTree4(p *tree.TreeNode, q *tree.TreeNode) bool {
	pStack := []*tree.TreeNode{p}
	qStack := []*tree.TreeNode{q}

	for len(pStack) > 0 && len(qStack) > 0 {
		pNode := pStack[len(pStack)-1]
		qNode := qStack[len(qStack)-1]

		pStack = pStack[:len(pStack)-1]
		qStack = qStack[:len(qStack)-1]

		if pNode == nil && qNode == nil {
			continue
		}

		if pNode == nil || qNode == nil {
			return false
		}

		if pNode.Val != qNode.Val {
			return false
		}

		pStack = append(pStack, pNode.Left, pNode.Right)
		qStack = append(qStack, qNode.Left, qNode.Right)
	}

	return len(pStack) == 0 && len(qStack) == 0
}

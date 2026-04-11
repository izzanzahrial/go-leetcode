package linkedlist

// https://leetcode.com/problems/leaf-similar-trees/description/?envType=study-plan-v2&envId=leetcode-75
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leafs1, leafs2 []int

	countLeaf(root1, &leafs1)
	countLeaf(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i, _ := range leafs1 {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}

func countLeaf(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
		return
	}

	countLeaf(root.Left, result)
	countLeaf(root.Right, result)
}

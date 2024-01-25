package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/maximum-binary-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, idx := maxVal(nums...)

	root := &TreeNode{
		Val:   max,
		Left:  ConstructMaximumBinaryTree(nums[:idx]),
		Right: ConstructMaximumBinaryTree(nums[idx+1:]),
	}

	return root
}

func maxVal(nums ...int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	max, idx := nums[0], 0

	for i, num := range nums {
		if num > max {
			max = num
			idx = i
		}
	}

	return max, idx
}

package tree

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/description/
type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := &[]int{}
	dfs(result, root)
	return *result
}

func dfs(result *[]int, node *Node) {
	*result = append(*result, node.Val)
	for _, child := range node.Children {
		dfs(result, child)
	}

	return
}

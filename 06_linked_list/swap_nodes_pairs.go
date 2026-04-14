package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/swap-nodes-in-pairs/submissions/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{}
	curr := head
	new := head.Next

	for curr != nil && curr.Next != nil {
		node1 := curr
		node2 := curr.Next

		// point the node 1 next to the next of node 2
		node1.Next = node2.Next

		// point the node 2 next to the node 1
		node2.Next = node1

		if prev != nil {
			prev.Next = node2
		}

		prev, curr = node1, node1.Next
	}

	return new
}

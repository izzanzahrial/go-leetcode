package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/reverse-nodes-in-k-group/submissions/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	l := lenNodes(head)

	if l < k {
		return head
	}

	result := &ListNode{Next: head}
	prev := result

	for l >= k {
		curr := prev.Next

		for i := 1; i < k; i++ {
			next := curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}

		prev = curr
		l -= k
	}

	return result.Next
}

func lenNodes(node *ListNode) int {
	var count int

	for node != nil {
		node = node.Next
		count++
	}

	return count
}

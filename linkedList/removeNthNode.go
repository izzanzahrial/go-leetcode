package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Still error for linked list that have length 1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	nextNode := slow.Next.Next
	nthNode := slow.Next
	slow.Next = nextNode
	nthNode.Next = nil

	return head
}

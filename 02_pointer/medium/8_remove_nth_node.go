package pointer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	firstPointer := head
	secondPointer := head

	for n > 0 {
		secondPointer = secondPointer.Next
		n--
	}

	if secondPointer == nil {
		return head.Next
	}

	for secondPointer.Next != nil {
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next
	}

	firstPointer.Next = firstPointer.Next.Next

	return head
}

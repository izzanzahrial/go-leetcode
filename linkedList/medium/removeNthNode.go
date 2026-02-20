package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/linkedList"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// the idea is by using two pointer where the second pointer
// will be in the range of the nth node to the end of the list\
// if the second pointer is nil, then the first pointer will be the at the nth node
func removeNthFromEnd(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	firstPointer := head
	secondPointer := head

	for n > 0 {
		secondPointer = secondPointer.Next
		n--
	}

	// if the second pointer is nil, then the first pointer will be the at the nth node
	// so we return without the first pointer value
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

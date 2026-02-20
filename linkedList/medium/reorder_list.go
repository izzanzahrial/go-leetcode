package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/linkedList"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode.com/problems/reorder-list/
func reorderList(head *linkedlist.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// search the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half
	secondHalf := slow.Next
	slow.Next = nil
	var prev *linkedlist.ListNode
	for secondHalf != nil {
		temp := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = temp
	}

	// merge the secondHalf with the firstHalf
	firstHalf := head
	for prev != nil {
		temp1 := firstHalf.Next
		temp2 := prev.Next
		firstHalf.Next = prev
		prev.Next = temp1
		firstHalf = temp1
		prev = temp2
	}
}

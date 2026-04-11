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

	// make the first half next nil
	slow.Next = nil
	var prev *linkedlist.ListNode
	for secondHalf != nil {
		nextNode := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = nextNode
	}

	// merge the secondHalf with the firstHalf
	secondHalf = prev
	firstHalf := head
	for secondHalf != nil {
		next1 := firstHalf.Next
		next2 := secondHalf.Next
		firstHalf.Next = secondHalf
		secondHalf.Next = next1
		firstHalf = next1
		secondHalf = next2
	}
}

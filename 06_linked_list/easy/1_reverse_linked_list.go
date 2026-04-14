package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/06_linked_list"

func reverseList(head *linkedlist.ListNode) *linkedlist.ListNode {
	var prevNode *linkedlist.ListNode

	for head != nil {
		// temporary node to hold the next node
		nextNode := head.Next

		// point the current node next to the previous node
		head.Next = prevNode

		// make the current node as the previous node
		prevNode = head

		// move the head to the next node
		head = nextNode
	}

	return prevNode
}

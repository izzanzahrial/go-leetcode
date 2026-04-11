package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/?envType=study-plan-v2&envId=leetcode-75
func deleteMiddle(head *ListNode) *ListNode {
	mid := findMid(head)
	if mid == 0 {
		return nil
	}

	prev := &ListNode{}
	curr := head
	next := head.Next

	for mid != 0 {
		prev = curr
		curr = curr.Next
		next = next.Next
		mid--
	}

	prev.Next = next
	return head
}

func findMid(node *ListNode) int {
	length := findLen(node)
	return length / 2
}

func findLen(node *ListNode) int {
	len := 1

	for node.Next != nil {
		len += 1
		node = node.Next
	}

	return len
}

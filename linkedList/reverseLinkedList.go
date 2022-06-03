package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode

	for head != nil {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		head = nextNode
	}

	return prevNode
}

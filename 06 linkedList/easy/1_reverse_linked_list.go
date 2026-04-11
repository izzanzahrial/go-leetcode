package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode

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

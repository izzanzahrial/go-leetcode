package linkedlist

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?isFullScreen=true
func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	head := llist

	for i := 0; i < int(position); i++ {
		if i == int(position)-1 {
			next := llist.next
			curr := &SinglyLinkedListNode{
				data: data,
				next: next,
			}
			llist.next = curr
		} else {
			llist = llist.next
		}
	}

	return head
}

package linkedlist

// https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem?isFullScreen=true
func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		// Deleting the head node
		if llist != nil { //check to see if the list is actually there
			return llist.next // Return the new head (which is the next node)
		} else {
			return nil // if list is empty we do nothing
		}

	}

	curr := llist
	for i := 0; i < int(position)-1; i++ {
		if curr == nil || curr.next == nil {
			// Position is out of bounds.  Returning the list as is.
			return llist
		}
		curr = curr.next
	}

	//check to see if the next pointer is nil, if so do nothing
	if curr != nil && curr.next != nil {
		curr.next = curr.next.next // perform the deletion
	}

	return llist
}

package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/linkedList"

// https://leetcode.com/problems/copy-list-with-random-pointer/
// the idea is by storing the value of the node in the map
// then after all the value is stored, point them one by one
// by the value of the current node
func copyRandomList(head *linkedlist.Node) *linkedlist.Node {
	if head == nil {
		return nil
	}

	oldToNew := make(map[*linkedlist.Node]*linkedlist.Node)

	// store the value of the node in the map
	curr := head
	for curr != nil {
		oldToNew[curr] = &linkedlist.Node{Val: curr.Val}
		curr = curr.Next
	}

	// point them one by one by the value of the current node
	curr = head
	for curr != nil {
		oldToNew[curr].Next = oldToNew[curr.Next]
		oldToNew[curr].Random = oldToNew[curr.Random]
		curr = curr.Next
	}

	// return the head of the new list
	return oldToNew[head]
}

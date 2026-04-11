package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/linkedList"

/**
 * Definition for singly-linked list.
 * type linkedlist.ListNode struct {
 *     Val int
 *     Next *linkedlist.ListNode
 * }
 */
// https://leetcode.com/problems/merge-k-sorted-lists/submissions/
func MergeKLists(lists []*linkedlist.ListNode) *linkedlist.ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwo(result, lists[i])
	}

	return result
}

func mergeTwo(l1, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	result := &linkedlist.ListNode{}
	curr := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}

	return result.Next
}

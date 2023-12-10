package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/merge-k-sorted-lists/submissions/
func MergeKLists(lists []*ListNode) *ListNode {
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

func mergeTwo(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
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

package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/06_linked_list"

/**
 * Definition for singly-linked list.
 * type linkedlist.ListNode struct {
 *     Val int
 *     Next *linkedlist.ListNode
 * }
 */
// https://leetcode.com/problems/merge-two-sorted-lists/description/
// func MergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
// 	result := &linkedlist.ListNode{}

// 	for curr := result; list1 != nil || list2 != nil; curr = curr.Next {
// 		if list1 == nil {
// 			curr.Next = list2
// 			break
// 		} else if list2 == nil {
// 			curr.Next = list1
// 			break
// 		} else if list1.Val < list2.Val {
// 			curr.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			curr.Next = list2
// 			list2 = list2.Next
// 		}
// 	}

// 	return result.Next
// }

/**
 * Definition for singly-linked list.
 * type linkedlist.ListNode struct {
 *     Val int
 *     Next *linkedlist.ListNode
 * }
 */
func MergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	result := &linkedlist.ListNode{}
	curr := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return result.Next
}

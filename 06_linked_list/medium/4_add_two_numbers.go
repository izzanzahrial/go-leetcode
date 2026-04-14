package linkedlist

import linkedlist "github.com/izzanzahrial/go-leetcode/linkedList"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	result := &linkedlist.ListNode{}
	prev := result
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// check if the carry to the next value is greater than 9
		// if it is, then carry will have value 1, e.g. sum = 10
		// meaning carry = 10 / 10 = 1
		carry = sum / 10
		// for more readablity use this
		// if sum >= 10 {
		//     carry = 1
		// } else {
		//     carry = 0
		// }

		prev.Next = &linkedlist.ListNode{
			// counting whats left after mod 10 incase it is greater than 9
			Val: sum % 10,
		}

		prev = prev.Next
	}

	return result.Next
}

package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	rabbit := head.Next
	turtle := head

	for rabbit != nil && rabbit.Next != nil {
		if rabbit == turtle {
			return true
		}

		rabbit = rabbit.Next.Next
		turtle = turtle.Next
	}

	return false
}

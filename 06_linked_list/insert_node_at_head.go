package linkedlist

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list/problem?isFullScreen=true
// reverse the num input
func InsertNodeAtHead() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing input:", err)
			return
		}
		nums = append(nums, int(value))
	}

	dummyHead := &SinglyLinkedListNode{}
	current := dummyHead
	for i := len(nums) - 1; i >= 0; i-- {
		node := &SinglyLinkedListNode{
			data: int32(nums[i]),
		}
		current.next = node
		current = current.next
	}

	for node := dummyHead.next; node != nil; node = node.next {
		fmt.Println(node.data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

// reverse the linked list
func InsertNodeAtHead2() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing input:", err)
			return
		}
		nums = append(nums, int(value))
	}

	dummyHead := &SinglyLinkedListNode{}
	current := dummyHead
	for i := 0; i < len(nums); i++ {
		node := &SinglyLinkedListNode{
			data: int32(nums[i]),
		}
		node.next = current
		current = node
	}

	// dummyHead.data = listOfllNums[len(listOfllNums)-1]
	for node := current; node.next != nil; node = node.next {
		fmt.Println(node.data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

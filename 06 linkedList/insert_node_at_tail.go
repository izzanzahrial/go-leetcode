package linkedlist

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list/problem?isFullScreen=true
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dummyHead := &SinglyLinkedListNode{}
	current := dummyHead
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		if lineNumber == 1 {
			continue
		}

		line := scanner.Text()
		value, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return
		}

		newNode := &SinglyLinkedListNode{
			data: int32(value),
		}
		current.next = newNode
		current = current.next
	}

	for node := dummyHead.next; node != nil; node = node.next {
		fmt.Println(node.data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

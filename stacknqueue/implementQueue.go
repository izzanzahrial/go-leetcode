package stacknqueue

type MyQueue struct {
	stack []int
}

func QueueConstructor() MyQueue {
	return MyQueue{stack: []int{}}
}

func (q *MyQueue) Push(x int) {
	q.stack = append(q.stack, x)
}

func (q *MyQueue) Pop() int {
	val := q.stack[0]
	q.stack = q.stack[1:]
	return val
}

func (q *MyQueue) Peek() int {
	return q.stack[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

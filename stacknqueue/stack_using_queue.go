package stacknqueue

// https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	stack []int
}

func StackConstructor() MyStack {
	return MyStack{stack: make([]int, 0)}
}

func (s *MyStack) Push(x int) {
	s.stack = append([]int{x}, s.stack[:]...)
}

func (s *MyStack) Pop() int {
	val := s.stack[0]
	s.stack = s.stack[1:]
	return val
}

func (s *MyStack) Top() int {
	return s.stack[0]
}

func (s *MyStack) Empty() bool {
	return len(s.stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

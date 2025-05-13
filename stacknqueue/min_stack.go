package stacknqueue

// https://leetcode.com/problems/min-stack/description/
// The idea of this solution is to create another stack to account for the current minimum value
type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push not only append to the stack
// but also append to the minStack if the value is less than the current min
// if the current min value is less than the value being pushed, then append the current min value
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

// Pop from both stacks, because if the value being popped is the current min value
// then we also need to remove it from the minStack, that's why we pop from both
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package datastructure

// https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	stack []int
}

func StackConstructor() MyStack {
	return MyStack{
		stack: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	num := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return num
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

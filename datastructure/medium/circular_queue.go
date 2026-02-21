package datastructure

// https://leetcode.com/problems/design-circular-queue/description/
type MyCircularQueue struct {
	queue  []int
	length int
}

func CircularQueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:  make([]int, 0, k),
		length: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.length {
		return false
	}

	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}

	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}

	return -1
}

func (this *MyCircularQueue) Rear() int {
	if len(this.queue) > 0 {
		return this.queue[len(this.queue)-1]
	}

	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.length
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */

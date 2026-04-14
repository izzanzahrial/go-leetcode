package stacknqueue

type FreqStack struct {
	freqMap map[int]int   // Maps a number to its current frequency
	group   map[int][]int // Maps a frequency to a list (stack) of numbers that have this frequency
	maxFreq int           // Tracks the current highest frequency
}

func NewFreqStack() FreqStack {
	return FreqStack{
		freqMap: make(map[int]int),
		group:   make(map[int][]int),
		maxFreq: 0,
	}
}

func (this *FreqStack) Push(val int) {
	// 1. Increment the frequency of val
	this.freqMap[val]++
	currentFreq := this.freqMap[val]

	// 2. Add val to the list of numbers that have this currentFreq.
	// The elements in group[currentFreq] are ordered by the time they achieved this frequency.
	// The last element is the one most recently achieving this frequency.
	this.group[currentFreq] = append(this.group[currentFreq], val)

	// 3. Update maxFreq if this new frequency is higher
	if currentFreq > this.maxFreq {
		this.maxFreq = currentFreq
	}
}

func (this *FreqStack) Pop() int {
	// 1. Get the list of numbers that are at the current maxFreq
	elementsAtMaxFreq := this.group[this.maxFreq]

	// 2. The number to pop is the last one added to this list
	// (this handles the "closest to stack's top" tie-breaker)
	valToPop := elementsAtMaxFreq[len(elementsAtMaxFreq)-1]

	// 3. Remove this number from the list for this.maxFreq
	this.group[this.maxFreq] = elementsAtMaxFreq[:len(elementsAtMaxFreq)-1]

	// 4. Decrement its frequency in freqMap
	this.freqMap[valToPop]--

	// 5. If the list for this.maxFreq is now empty, it means no numbers
	// currently have this frequency. So, the maxFreq might need to decrease.
	// This happens if valToPop was the last element having this.maxFreq.
	// (It's guaranteed that pop is only called if the stack is not empty)
	if len(this.group[this.maxFreq]) == 0 {
		this.maxFreq--
		// We don't need to delete this.group[this.maxFreq] entry,
		// it will just be an empty slice, which is fine.
		// Or, `delete(this.group, this.maxFreq)` if this.maxFreq became 0 due to this pop.
		// But simply decrementing maxFreq is sufficient.
	}

	return valToPop
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

type FreqStack2 struct {
	stack       []int
	freqMap     map[int][]int // array of idx
	freq        int
	mostFreq    int
	mostFreqIdx int
}

func Constructor() FreqStack2 {
	return FreqStack2{
		stack:   []int{},
		freqMap: make(map[int][]int),
	}
}

func (this *FreqStack2) Push(val int) {
	this.stack = append(this.stack, val)
	this.freqMap[val] = append(this.freqMap[val], len(this.stack)-1)
	this.freq, this.mostFreq, this.mostFreqIdx = this.calculateMostFreq()
}

func (this *FreqStack2) Pop() int {
	val := this.mostFreq
	removeIdx := this.mostFreqIdx

	// remove from stack
	this.stack = append(this.stack[:removeIdx], this.stack[removeIdx+1:]...)

	// remove last occurrence index for that value
	if len(this.freqMap[val]) > 0 {
		this.freqMap[val] = this.freqMap[val][:len(this.freqMap[val])-1]
	}
	this.freq, this.mostFreq, this.mostFreqIdx = this.calculateMostFreq()

	// shift all indices > removed index
	for k, arr := range this.freqMap {
		for i := range arr {
			if arr[i] > removeIdx {
				arr[i]--
			}
		}
		this.freqMap[k] = arr
	}

	this.freq, this.mostFreq, this.mostFreqIdx = this.calculateMostFreq()
	return val
}

func (this *FreqStack2) calculateMostFreq() (freq int, mostFreq int, mostFreqIdx int) {
	for num, node := range this.freqMap {
		if len(node) == 0 {
			continue
		}

		// calculate the frequency of the num by the lenght of indexes in the map
		if len(node) > freq {
			mostFreq = num
			// the most recent idx of the num is the last idx in the node map
			mostFreqIdx = node[len(node)-1]
			freq = len(node)
			continue
		}

		if len(node) == freq {
			if node[len(node)-1] < mostFreqIdx {
				continue
			}

			mostFreq = num
			mostFreqIdx = node[len(node)-1]
		}
	}

	return freq, mostFreq, mostFreqIdx
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */

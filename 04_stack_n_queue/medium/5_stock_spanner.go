package stacknqueue

// https://leetcode.com/problems/online-stock-span/description/
type StockSpanner struct {
	stockStack [][]int // price, span
}

func Constructor() StockSpanner {
	return StockSpanner{
		stockStack: make([][]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	// if the stack is empty we skip
	// and keep looping if the prev price in the stack is lower than the current price
	for len(this.stockStack) > 0 && this.stockStack[len(this.stockStack)-1][0] <= price {
		span += this.stockStack[len(this.stockStack)-1][1]
		this.stockStack = this.stockStack[:len(this.stockStack)-1]
	}

	this.stockStack = append(this.stockStack, []int{price, span})
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */

type StockSpanner2 struct {
	stack []int
}

func StockSpannerConstructor() StockSpanner2 {
	return StockSpanner2{
		stack: []int{},
	}
}

func (this *StockSpanner2) Next(price int) int {
	stockSpan := 1
	for i := len(this.stack) - 1; i >= 0; i-- {
		if price < this.stack[i] {
			break
		}

		if price >= this.stack[i] {
			stockSpan++
		}
	}

	this.stack = append(this.stack, price)

	return stockSpan
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

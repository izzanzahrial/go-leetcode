package datastructure

// https://leetcode.com/problems/online-stock-span/description/
type StockSpanner struct {
	stack []int
}

func StockSpannerConstructor() StockSpanner {
	return StockSpanner{
		stack: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
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

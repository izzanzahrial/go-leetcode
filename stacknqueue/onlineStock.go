package stacknqueue

// https://leetcode.com/problems/online-stock-span/?envType=study-plan-v2&envId=leetcode-75
type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	result := 1
	this.stack = append(this.stack, price)

	if len(this.stack) == 1 {
		return result
	}

	for i := len(this.stack) - 2; i >= 0; i-- {
		if this.stack[i] > price {
			return result
		}
		result++
	}

	return result
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

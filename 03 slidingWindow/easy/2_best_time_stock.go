package slidingwindow

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	lowestPrice, highestProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		currPrice := prices[i]
		if currPrice < lowestPrice {
			lowestPrice = currPrice
		}

		if currPrice > lowestPrice {
			currProfit := currPrice - lowestPrice
			highestProfit = max(highestProfit, currProfit)
		}
	}

	return highestProfit
}

func maxProfit2(prices []int) int {
	maxProfit := 0
	pivot := math.MaxInt
	for _, price := range prices {
		pivot = min(pivot, price)
		maxProfit = max(maxProfit, price-pivot)
	}

	return maxProfit
}

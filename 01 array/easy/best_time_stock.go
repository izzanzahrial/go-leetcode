package array

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

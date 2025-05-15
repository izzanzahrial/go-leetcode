package binarysearch

// the idea is creating left and right bound
// the left bound is the minimal capactiy that the ship need to carry for all the weights
// meaning, the heaviest weight within the packages
// the right bound is the maximal capactiy that the ship need to carry for all the weights
// meaning all the sum of weights
// using left capacity, the ship at maximum can deliver the package within length of weights days
// using right capacity, the ship at minimum can deliver the package within 1 day, but using max capacity
func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, num := range weights {
		if num > left {
			left = num
		}

		right += num
	}

	isShipable := func(cap int) bool {
		currDays, currCap := 1, cap
		for _, w := range weights {
			if currCap-w < 0 {
				currDays++
				currCap = cap
			}

			currCap -= w
		}

		return currDays <= days
	}

	result := right
	for left <= right {
		mid := left + (right-left)/2
		ok := isShipable(mid)
		if ok {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

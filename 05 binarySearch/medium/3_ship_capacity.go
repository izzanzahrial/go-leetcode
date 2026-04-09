package binarysearch

// https://neetcode.io/problems/capacity-to-ship-packages-within-d-days/question
func shipWithinDays(weights []int, days int) int {
	// the idea is creating left and right bound
	// the left bound is the minimal capactiy that the ship need to carry for all the weights
	// meaning, the heaviest weight within the packages
	// the right bound is the maximal capactiy that the ship need to carry for all the weights
	// meaning all the sum of weights
	var minWeight, maxWeight int
	for _, w := range weights {
		if w > minWeight {
			minWeight = w
		}

		maxWeight += w
	}

	isShipAble := func(capacity int) bool {
		currDay := 1
		currWeight := capacity
		for _, w := range weights {
			// if the current weight minus the next weight that gonna be ship
			// is less than 0, meaning we cannot ship it
			// we add the day + 1
			// return the current weight to original capacity
			if currWeight-w < 0 {
				currDay += 1
				currWeight = capacity
			}

			// then we minus it again with the next weight
			currWeight = currWeight - w
		}

		return currDay <= days
	}

	var result int
	for minWeight <= maxWeight {
		midWeight := minWeight + ((maxWeight - minWeight) / 2)

		// if ship able, we still need to check min capacity
		// meaning we still need to check is lower value is possible
		// but we record the current value as a result
		if isShipAble(midWeight) {
			maxWeight = midWeight - 1
			result = midWeight
		} else {
			minWeight = midWeight + 1
		}
	}

	return result
}

func shipWithinDays2(weights []int, days int) int {
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

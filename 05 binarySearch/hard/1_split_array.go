package binarysearch

// the idea of this solution is we need to search the lowest and the highest number
// for any possible sum of subarray. For the lowest number we need the higest number in nums array
// so we can create a subarray for each number in nums array, meaning we will have subarray with the length of nums array.
// For the highest number we need the total number in nums array, so we can create a subbaray that have the total sum of
// the entire nums array
//
// After that test the middle number between the lowest and the highest number possible to split the nums array
// if it's possible, we move the highest number into the middle number - 1, to try get the minimal number of sum
// if it's not possible, we have add the middle number, so we move the lowest number to the middle number + 1
func splitArray(nums []int, k int) int {
	var low, high int
	for _, num := range nums {
		// highest number in nums
		low = max(low, num)

		// total number in nums
		high += num
	}

	splitable := func(sum int) bool {
		// the subArrays have value of 1, because at least minimal
		// we can create 1 subArrays which is the entire array
		currSum, subArrays := 0, 1
		for _, num := range nums {
			currSum += num
			if currSum > sum {
				subArrays++
				currSum = num

				// early exit if subArrays > the maximal number of subarray that we want (k)
				if subArrays > k {
					return false
				}
			}
		}

		return true
	}

	for low <= high {
		currSum := low + (high-low)/2
		// if splitable meaning we can split it by the number of subarray of k
		// we have to test the minimum again by moving the high number into currSum(mid) - 1
		if splitable(currSum) {
			high = currSum - 1
		} else {
			// not splitable meaning we need higher number for the current sum
			// we move the low number into the currSum(mid) + 1
			low = currSum + 1
		}
	}

	return low
}

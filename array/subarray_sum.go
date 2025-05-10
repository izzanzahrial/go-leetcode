package array

// https://leetcode.com/problems/subarray-sum-equals-k/description/
// Brute force
func subarraySum(nums []int, k int) int {
	count := 0
	n := len(nums)

	// Iterate over all possible start points of a subarray
	for i := 0; i < n; i++ {
		currentSum := 0
		// Iterate over all possible end points for the current start point
		for j := i; j < n; j++ {
			currentSum += nums[j] // Add current element to sum
			if currentSum == k {
				count++
			}
		}
	}

	return count
}

package slidingwindow

// TODO: need some fixing, missing result by 1
// https://leetcode.com/problems/max-consecutive-ones-iii/solutions/3477843/golang-solution-good-explanation/?envType=study-plan-v2&envId=leetcode-75
func longestOnes(nums []int, k int) int {
	var max int
	var left, right int

	for i, _ := range nums {
		var curr int
		left = i
		for curr <= k && i < len(nums) {
			if nums[i] == 0 {
				curr++
			}

			right = i
			i++
		}

		result := right - left + 1
		if max < result {
			max = result
		}
	}

	return max
}

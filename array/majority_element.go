package array

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	currNum := nums[0]
	currResult := 1

	// start at index 1 because we already initialize the first num in currNum
	for i := 1; i < len(nums); i++ {
		if currNum != nums[i] {
			currResult--

			// if the current result equal to 0, we change the currNum
			if currResult == 0 {
				currResult = 1
				currNum = nums[i]
			}
		} else {
			currResult++
		}
	}

	return currNum
}

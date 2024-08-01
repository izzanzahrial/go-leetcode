package hash

// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		numMap[num] = true
	}

	var result int
	for _, num := range nums {
		if numMap[num+1] {
			continue
		}

		var curr int
		for i := num; numMap[i]; i-- {
			curr++
		}

		if curr > result {
			result = curr
		}
	}

	return result
}

package hash

// https://leetcode.com/problems/longest-consecutive-sequence/
// search descending (backward from the biggest to the smallest number)
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

// search ascending (forward from smallest number to the biggest number)
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	var longest int
	for num := range numsMap {
		// search for the start of the number, if the current number - 1 exists
		// meaning current number is not the start of the number
		if numsMap[num-1] {
			continue
		}

		currNum := num
		currLength := 1

		for numsMap[currNum+1] {
			currNum++
			currLength++
		}

		if currLength > longest {
			longest = currLength
		}
	}

	return longest
}

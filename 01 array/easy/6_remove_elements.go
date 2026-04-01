package array

// https://leetcode.com/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			k++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			// why we subtract the i index, because after we delete the current num index
			// the index i will move forward, the previous next index will be at the current num index
			// it will get skip, in order for it to not get skipped, we have to make the i index stay in the last index
			// meaning substract the index
			i--
		}
	}

	return k
}

func removeElement2(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

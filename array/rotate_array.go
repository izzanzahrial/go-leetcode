package array

// in place solution
// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	numsLength := len(nums)
	// check if is looping, only account for the actual rotation
	// whats the left amount of rotation modulus by the length of nums
	modK := k % numsLength
	// edge case if the rotation still 0 is early return
	if modK == 0 {
		return
	}

	// create temporary array for num that get rotate to the front
	temp := make([]int, 0, modK)
	idx := 0
	lastIdx := numsLength - 1
	for idx != modK {
		temp = append(temp, nums[lastIdx-idx])
		idx++
	}

	// move the num that doesn't get rotated to the back
	for i := lastIdx - modK; i >= 0; i-- {
		nums[lastIdx] = nums[i]
		lastIdx--
	}

	// add the num from temp to the actual num array
	lastIdx = modK - 1
	for _, num := range temp {
		nums[lastIdx] = num
		lastIdx--
	}

	return
}

// Copy solution
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[(i+k)%n] = nums[i]
	}
	copy(nums, res)
}

func rotate3(nums []int, k int) {
	numsLen := len(nums)
	modK := k % numsLen
	if modK == 0 {
		return
	}

	tempNums := append([]int{}, nums[numsLen-modK:]...)
	lastIdx := len(nums) - 1
	for i := lastIdx - modK; i >= 0; i-- {
		nums[lastIdx] = nums[i]
		lastIdx--
	}

	for i, num := range tempNums {
		nums[i] = num
	}
}

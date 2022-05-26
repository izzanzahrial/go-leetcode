package main

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var curr int

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]*nums[i-1] {
			if nums[i] > curr {
				curr = nums[i]
			}
		} else {
			nums[i] = nums[i-1] * nums[i]
			if nums[i] > curr {
				curr = nums[i]
			}
		}
	}
	return curr
}

package array

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	combine := append(nums1, nums2...)

	sort.Ints(combine)

	val1 := len(combine) / 2

	if len(combine)%2 == 0 {
		val2 := val1 - 1
		result := combine[val1] + combine[val2]
		return float64(result) / 2
	}

	return float64(combine[val1])
}

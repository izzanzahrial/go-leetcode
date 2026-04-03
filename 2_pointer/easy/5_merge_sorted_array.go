package pointer

// https://leetcode.com/problems/merge-sorted-array/description/
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]

		currIdx := m + i
		for j := currIdx - 1; j >= 0; j-- {
			if nums1[currIdx] < nums1[j] {
				nums1[currIdx], nums1[j] = nums1[j], nums1[currIdx]
				currIdx--
			}
		}
	}
}

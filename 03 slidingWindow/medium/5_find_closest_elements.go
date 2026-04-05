package slidingwindow

// https://leetcode.com/problems/find-k-closest-elements/
func findClosestElements(arr []int, k int, x int) []int {
	for len(arr) > k {
		absA := abs(arr[0] - x)
		absB := abs(arr[len(arr)-1] - x)
		if absA > absB {
			arr = arr[1:]
		} else {
			arr = arr[:len(arr)-1]
		}
	}

	return arr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

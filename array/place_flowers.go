package array

// https://leetcode.com/problems/can-place-flowers/description/?envType=study-plan-v2&envId=leetcode-75
func CanPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	total := 0
	for i := 0; i < m; i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == m-1 || flowerbed[i+1] == 0) {
				total++
				flowerbed[i] = 1
			}
		}
	}
	return total >= n
}

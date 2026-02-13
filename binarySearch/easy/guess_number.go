package binarysearch

import "fmt"

// https://leetcode.com/problems/guess-number-higher-or-lower/description/
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		num := guess(mid)
		fmt.Println(mid, num)
		if num == 0 {
			return mid
		}

		if num == -1 {
			right = mid - 1
		}

		if num == 1 {
			left = mid + 1
		}
	}

	return 0
}

func guess(num int) int {
	return 0
}

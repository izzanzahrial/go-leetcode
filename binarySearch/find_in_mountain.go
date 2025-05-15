package binarysearch

import "fmt"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
type MountainArray struct {
}

func (this *MountainArray) get(index int) int {
	return 0
}

func (this *MountainArray) length() int {
	return 0
}

// https://leetcode.com/problems/find-in-mountain-array/description/
func findInMountainArray(target int, mountainArr *MountainArray) int {
	arrLength := mountainArr.length()
	leftIdx, rightIdx := 0, arrLength-1
	var peakIdx int
	for leftIdx < rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		midNum := mountainArr.get(midIdx)
		rightNum := mountainArr.get(midIdx + 1)

		if midNum < rightNum {
			// Ascending slope, peak is to the right
			leftIdx = midIdx + 1
		} else {
			// Descending slope or peak itself, peak is 'mid' or to its left
			rightIdx = midIdx
		}
	}
	peakIdx = leftIdx // When loop terminates, peakLeft == peakRight, and this is the peak index.

	// since we only need the first index of the target
	// we should calculate the left side of the peak first
	// so if we find it, we don't have to calculate the right side
	// we can early return it
	leftIdx = 0
	rightIdx = peakIdx
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		midNum := mountainArr.get(midIdx)
		fmt.Println(leftIdx, rightIdx, midIdx, midNum)
		if midNum == target {
			return midIdx
		}

		if midNum < target {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	leftIdx = peakIdx
	rightIdx = arrLength - 1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		midNum := mountainArr.get(midIdx)
		fmt.Println(leftIdx, rightIdx, midIdx, midNum)
		if midNum == target {
			return midIdx
		}

		// switch the index calculation since we are on the right side of the mountain
		// meaning the more right index we get, the less the number
		// so if the middle number is less then target, we have to move the right index closer to the middle
		// if the middle number is bigger than target, move the left index far away from the middle
		if midNum < target {
			rightIdx = midIdx - 1
		} else {
			leftIdx = midIdx + 1
		}
	}

	return -1
}

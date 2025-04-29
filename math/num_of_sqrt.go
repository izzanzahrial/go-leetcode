package math

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=true
func Squares(a int32, b int32) int32 {
	// Handle edge cases for non-negative input ranges
	if a < 0 {
		fmt.Println("Warning: min must be non-negative for real square roots.")
		// Decide how to handle: return error, panic, or just return empty
		return 0
	}
	if a > b {
		return 0 // Return empty slice for invalid range
	}

	// 1. Calculate the smallest integer s such that s*s >= min
	//    We take the square root of min, and round up to the nearest integer.
	startSqrtFloat := math.Sqrt(float64(a))
	startSqrtInt := int(math.Ceil(startSqrtFloat))

	// 2. Calculate the largest integer s such that s*s <= max
	//    We take the square root of max, and round down to the nearest integer.
	endSqrtFloat := math.Sqrt(float64(b))
	endSqrtInt := int(math.Floor(endSqrtFloat))

	// If the calculated start is greater than the calculated end,
	// it means there are no perfect squares in the range [min, max].
	if startSqrtInt > endSqrtInt {
		return 0
	}

	results := make([]int, 0, endSqrtInt-startSqrtInt+1)

	// 4. Iterate through the potential integer square roots
	//    Any integer s in this range [startSqrtInt, endSqrtInt]
	//    will have its square s*s within the range [min, max].
	for s := startSqrtInt; s <= endSqrtInt; s++ {
		results = append(results, s)
	}

	return int32(len(results))
}

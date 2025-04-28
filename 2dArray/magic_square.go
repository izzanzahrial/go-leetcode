package twodarray

import "math"

// IMPORTANT!!! =========https://en.wikipedia.org/wiki/Magic_square=========
// https://www.hackerrank.com/challenges/magic-square-forming/problem?isFullScreen=true
func FormingMagicSquare(s [][]int32) int32 {
	// Define all 8 possible 3x3 magic squares using numbers 1-9
	// Each element in magicSquares is a [3][3]int32 array
	magicSquares := [][3][3]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	minCost := int32(math.MaxInt32)
	for _, square := range magicSquares {
		var currCost int32
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s[0]); j++ {
				currCost += int32(math.Abs(float64(square[i][j] - s[i][j])))
			}
		}

		if currCost < minCost {
			minCost = currCost
		}
	}

	return minCost
}

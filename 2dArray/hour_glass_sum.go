package twodarray

import "math"

// https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true
func hourglassSum(arr [][]int32) int32 {
	maxResult := int32(math.MinInt32)

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[0])-2; j++ {
			// sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
			// 	arr[i+1][j+1] +
			// 	arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			sum := currHourGlass(i, j, arr)
			maxResult = max(maxResult, sum)
		}
	}

	return maxResult
}

func currHourGlass(i, j int, arr [][]int32) int32 {
	var result int32
	var isMiddle bool

	for k := i; k < i+3; k++ {
	loop:
		for l := j; l < j+3; l++ {
			if isMiddle {
				result += arr[k][l+1]
				isMiddle = false
				break loop
			} else {
				result += arr[k][l]
				if l == j+2 {
					isMiddle = true
				}
			}
		}
	}

	return result
}

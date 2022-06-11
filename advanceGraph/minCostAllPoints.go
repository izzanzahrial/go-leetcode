package advancegraph

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/min-cost-to-connect-all-points/
// still wrong
func minCostConnectPoints(points [][]int) int {
	var result int
	hashMin := make(map[[2]int]int)

	for i := 0; i < len(points); i++ {
		currPoint := [2]int{points[i][0], points[i][1]}
		hashMin[currPoint] = math.MaxInt64
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			currPoint := [2]int{points[i][0], points[i][1]}
			nextPoint := [2]int{points[j][0], points[j][1]}
			currVal := int(math.Abs(float64(currPoint[0]-nextPoint[0]) + float64(currPoint[1]-nextPoint[1])))
			if currVal < hashMin[currPoint] {
				hashMin[currPoint] = currVal
			}
		}
	}

	fmt.Println(hashMin)

	for _, val := range hashMin {
		if val == 9223372036854775807 {
			val = 0
		}
		result += val
	}

	return result
}

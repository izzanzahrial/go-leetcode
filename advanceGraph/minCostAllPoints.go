package advancegraph

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/min-cost-to-connect-all-points/
// still wrong the manhattan algorithm
func minCostConnectPoints(points [][]int) int {
	var result int
	pointsMap := make(map[[2]int]int)

	for _, point := range points {
		currPoint := [2]int{point[0], point[1]}
		pointsMap[currPoint] = math.MaxInt64
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			currPoint := [2]int{points[i][0], points[i][1]}
			nextPoint := [2]int{points[j][0], points[j][1]}
			// this part, manhattan algorithm
			currVal := int(math.Abs(float64(currPoint[0]-nextPoint[0])) + math.Abs(float64(currPoint[1]-nextPoint[0])))
			if currVal < pointsMap[currPoint] {
				pointsMap[currPoint] = currVal
			}
		}
	}

	for key, val := range pointsMap {
		fmt.Println(key, val)
		result += val
	}

	return result
}

package sorting

import "sort"

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/?envType=study-plan-v2&envId=leetcode-75
func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	if len(points) == 1 {
		return 1
	}

	sortBallons(points)
	result := 1
	currY := points[0][1]

	for i := 1; i < len(points); i++ {
		if currY < points[i][0] {
			currY = points[i][1]
			result++
		}
	}

	return result
}

func sortBallons(balloons [][]int) {
	sort.Slice(balloons, func(i, j int) bool {
		return balloons[i][1] < balloons[j][1]
	})
}

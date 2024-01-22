package sorting

import "sort"

// https://leetcode.com/problems/non-overlapping-intervals/?envType=study-plan-v2&envId=leetcode-75
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sortIntervals(intervals)

	var nonOverlap int
	curr := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if curr[1] <= intervals[i][0] {
			curr = intervals[i]
			nonOverlap++
		}
	}

	return len(intervals) - nonOverlap - 1
}

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
}

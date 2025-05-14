package stacknqueue

import "sort"

// https://leetcode.com/problems/car-fleet/description/
// the idea of this solution is whetere the arrival time of car behind is faster than the car in front
// meaning the car behind can catch up to the car in front
func carFleet(target int, position []int, speed []int) int {
	var stack [][]int

	for i, pos := range position {
		stack = append(stack, []int{pos, speed[i]})
	}

	// sort the cars by their position
	sort.Slice(stack, func(i, j int) bool {
		return stack[i][0] < stack[j][0]
	})

	for i := len(stack) - 1; i >= 1; i-- {
		// get the time by dividing the distance by the speed
		arrivalTime := float64(target-stack[i][0]) / float64(stack[i][1])
		arrivalTime2 := float64(target-stack[i-1][0]) / float64(stack[i-1][1])
		// the arrival time of the car in front is greater meaning it's slower
		if arrivalTime >= arrivalTime2 {
			stack = append(stack[:i-1], stack[i:]...)
		}
	}

	return len(stack)
}

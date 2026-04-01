package array

// the key to understanding this solution lies in recognizing that it doesn't directly update the array with each query
// Instead, it employs a clever trick called the difference array approach. This technique significantly optimizes
// the process of applying multiple range updates.
//
// Explanation
//
// Building the Difference Array:
//
// twodArray[firstIdx] += val
// if lastIdx < n {
//      twodArray[lastIdx] -= val
// }
//
// twodArray[firstIdx] += val: This line adds val to the element at index firstIdx in twodArray.
// Crucially, think of this as marking the start of the range where val needs to be added to the original array.
//
// if lastIdx < n { twodArray[lastIdx] -= val }: This is where the magic happens. It subtracts val from the element at index lastIdx in twodArray.
// This marks the end of the range where val needs to be added. The lastIdx < n check is important to ensure we don't write out of bounds of the twodArray.
// If lastIdx is n, then the update continues until the end of the array, so no subtraction is needed.
//
// Why does this work?
// Imagine twodArray representing the differences between consecutive elements in the original array.
// Adding val at firstIdx means the original array increases by val from that point onward.
// Subtracting val at lastIdx means the original array stops increasing by val at that point.
//
// Reconstructing the Original Array and Finding the Maximum:
//
// var current int64
// for _, value := range twodArray {
//      current += value
//      if current > result {
//          result = current
//      }
// }
//
// current += value: This line iteratively calculates the actual values in the array.
// current acts as an accumulator. By summing the values in twodArray, we effectively "propagate" the changes across the array.
// current is essentially building up the original array's values.
//
// if current > result { result = current }: This line keeps track of the maximum value encountered so far.
// As we reconstruct the original array by summing the differences in twodArray,
// we continuously check if the current value (current) is greater than the current maximum (result).

// https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true
func arrayManipulation(n int32, queries [][]int32) int64 {
	twodArray := make([]int64, n+1)
	var result int64

	for _, query := range queries {
		firstIdx := query[0] - 1
		lastIdx := query[1]
		val := int64(query[2])

		twodArray[firstIdx] += val
		if lastIdx < n {
			twodArray[lastIdx] -= val
		}
	}

	var current int64
	for _, value := range twodArray {
		current += value
		if current > result {
			result = current
		}
	}

	return result
}

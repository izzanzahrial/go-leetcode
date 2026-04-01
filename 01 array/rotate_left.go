package array

// https://www.hackerrank.com/challenges/array-left-rotation/problem?isFullScreen=true
func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	var result []int32

	result = append(result, arr[d:]...)
	result = append(result, arr[:d]...)

	return result
}

func rotateLeft2(d int32, input []int32) []int32 {
	rotated := make([]int32, len(input))
	copy(rotated, input[d:])
	copy(rotated[len(input)-int(d):], input[:d])

	return rotated
}

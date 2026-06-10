package array

// https://leetcode.com/problems/rotate-image/description/
// https://www.youtube.com/watch?v=fMSJSS7eO1w&ab_channel=NeetCode
func Rotate(matrix [][]int) {
	left, right := 0, len(matrix)-1

	for left < right {
		for i := 0; i < right-left; i++ {
			top, bottom := left, right

			//hold first value
			topLeft := matrix[top][left+i]

			//counter clock wise

			//bottom left to top left
			matrix[top][left+i] = matrix[bottom-i][left]

			//bottom right to bottom left
			matrix[bottom-i][left] = matrix[bottom][right-i]

			//top right to bottom right
			matrix[bottom][right-i] = matrix[top+i][right]

			//top left to top right
			matrix[top+i][right] = topLeft
		}

		left += 1
		right -= 1
	}
}

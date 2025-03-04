package dynamicprogramming

// https://leetcode.com/problems/pascals-triangle/?envType=problem-list-v2&envId=dynamic-programming&
func generate(numRows int) [][]int {
	result := [][]int{{1}}
	if numRows == 1 {
		return result
	}

	for i := 2; i <= numRows; i++ {
		// create current result with the length of current pascal triangle
		// start at 2 because the first element is always with length of 1
		curr := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				curr[j] = 1
			} else {
				// the current result is the sum of the previous result at -1 index
				// and the previous result at the same index
				curr[j] = result[i-2][j-1] + result[i-2][j]
			}
		}
		result = append(result, curr)
	}

	return result
}

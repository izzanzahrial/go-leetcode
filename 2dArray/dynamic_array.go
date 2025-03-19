package twodarray

// https://www.hackerrank.com/challenges/dynamic-array/problem?isFullScreen=true
func dynamicArray(n int32, queries [][]int32) []int32 {
	arr := make([][]int32, n) // Create the 2D array
	for i := 0; i < int(n); i++ {
		arr[i] = make([]int32, 0) // Initialize each inner list as empty
	}

	var lastAnswer int32
	answers := make([]int32, 0)

	for _, query := range queries {
		queryType := query[0]
		x := query[1]
		y := query[2]

		idx := (x ^ lastAnswer) % n

		if queryType == 1 {
			arr[idx] = append(arr[idx], y)
		} else if queryType == 2 {
			lastAnswer = arr[idx][y%int32(len(arr[idx]))]
			answers = append(answers, lastAnswer)
		}
	}

	return answers
}

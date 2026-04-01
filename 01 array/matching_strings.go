package array

// https://www.hackerrank.com/challenges/sparse-arrays/problem?isFullScreen=true
func matchingStrings(stringList []string, queries []string) []int32 {
	mapQuery := make(map[string]int, len(queries))
	for _, query := range queries {
		mapQuery[query] = 0
	}

	for _, str := range stringList {
		if _, ok := mapQuery[str]; ok {
			mapQuery[str] += 1
		}
	}

	var result []int32
	for _, query := range queries {
		result = append(result, int32(mapQuery[query]))
	}

	return result
}

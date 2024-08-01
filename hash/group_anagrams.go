package hash

import "sort"

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)

	for _, str := range strs {
		sortStr := []rune(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})

		hashMap[string(sortStr)] = append(hashMap[string(sortStr)], str)
	}

	var result [][]string
	for i := range hashMap {
		var curr []string
		for _, str := range hashMap[i] {
			curr = append(curr, str)
		}
		result = append(result, curr)
	}

	return result
}

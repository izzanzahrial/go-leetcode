package array

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i, val := range haystack {
		if string(val) == string(needle[0]) {
			var count int

			if len(haystack)-i >= len(needle) {
				for j, v := range needle {
					if string(haystack[i+j]) != string(v) {
						break
					}

					count++
				}

				if count == len(needle) {
					return i
				}
			}
		}
	}

	return -1
}

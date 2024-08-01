package hash

// https://leetcode.com/problems/ransom-note/description/
func canConstruct(ransomNote string, magazine string) bool {
	ransomHash := make(map[rune]int)

	for _, run := range ransomNote {
		ransomHash[run]++
	}

	for _, run := range magazine {
		val, ok := ransomHash[run]
		if ok && val > 0 {
			ransomHash[run]--
		}
	}

	for _, val := range ransomHash {
		if val != 0 {
			return false
		}
	}

	return true
}

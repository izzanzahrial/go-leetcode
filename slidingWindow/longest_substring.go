package slidingwindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[rune]int)
	firstIdx, longest := 0, 0
	for idx, char := range s {
		// Check if the character exists in our current window
		if prevIdx, exists := idxMap[char]; exists && prevIdx >= firstIdx {
			// Move firstIdx to the position after the previous occurrence
			firstIdx = prevIdx + 1
		}

		idxMap[char] = idx
		currentLen := idx - firstIdx + 1
		if currentLen > longest {
			longest = currentLen
		}
	}

	return longest
}

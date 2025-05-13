package slidingwindow

import (
	"math"
	"reflect"
)

// https://leetcode.com/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {
	var minString string
	idxMap := make(map[rune][]int, len(t))
	subMap := make(map[rune]int, len(t))
	tMap := make(map[rune]int, len(t))
	for _, char := range t {
		tMap[char]++
	}

	for idx, char := range s {
		// only add to the subMap if the characters is in tMap
		if _, exists := tMap[char]; exists {
			subMap[char]++
			idxMap[char] = append(idxMap[char], idx)
		}

		// if the subMap has more characters than tMap
		if subMap[char] > tMap[char] {
			subMap[char]--
			var minIdx int
			for arrIdx, subIdx := range idxMap[char] {
				if subIdx < minIdx {
					minIdx = arrIdx
				}
			}

			temp := idxMap[char][:minIdx]
			idxMap[char] = append(temp, idxMap[char][minIdx+1:]...)
		}

		// check if the length submap is equal to the length of t
		// meaning we already get substring that have every character of t
		// we create the minimum substring from that
		if reflect.DeepEqual(subMap, tMap) {

			// we search the minimum and maximum idx
			// the minkey will be use to subtract the minimum idx after creating the minString
			// the minArrIdx will be use to subtract the minimum idx inside the idxMap array
			var minArrIdx int
			var minKey rune
			minIdx, maxIdx := math.MaxInt, math.MinInt
			for key, subArry := range idxMap {
				for arrIdx, subIdx := range subArry {
					maxIdx = max(maxIdx, subIdx)
					// we use this rather than the min function
					// because we also need to uptdate the minKey
					if subIdx < minIdx {
						minIdx = subIdx
						minKey = key
						minArrIdx = arrIdx
					}
				}
			}

			// if the current window is smaller than the current minString
			// we update the minString
			currLength := maxIdx - minIdx + 1
			if currLength < len(minString) || len(minString) == 0 {
				minString = s[minIdx : maxIdx+1]
			}

			// delete the minimum idx inside the idxMap array
			temp := idxMap[minKey][:minArrIdx]
			idxMap[minKey] = append(temp, idxMap[minKey][minArrIdx+1:]...)

			// subtract the minimum idx
			subMap[minKey]--
		}
	}

	return minString
}

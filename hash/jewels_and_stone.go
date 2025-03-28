package hash

// https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[rune]bool)

	for _, jewel := range jewels {
		jewelsMap[jewel] = true
	}

	var result int
	for _, stone := range stones {
		if _, ok := jewelsMap[stone]; ok {
			result++
		}
	}

	return result
}

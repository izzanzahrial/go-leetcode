package hash

func UniqueOccurrences(arr []int) bool {
	hashMap := make(map[int]int)
	isUnique := make(map[int]bool)

	for _, val := range arr {
		_, ok := hashMap[val]
		if ok {
			hashMap[val] += 1
		} else {
			hashMap[val] = 1
		}
	}

	for _, val := range hashMap {
		if ok := isUnique[val]; ok {
			return false
		}

		isUnique[val] = true
	}

	return true
}

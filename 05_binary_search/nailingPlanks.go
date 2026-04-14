package binarysearch

func nailingPlanks(plankA, plankB, nails []int) int {
	var minNail int

	for i := 0; i < len(plankA); i++ {
		currA := plankA[i]
		currB := plankB[i]
		if currA <= nails[i] && nails[i] <= currB {
			minNail = i + 1
		}
	}

	return minNail
}

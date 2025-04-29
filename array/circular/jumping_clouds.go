package circular

// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem?isFullScreen=true
func JumpingOnClouds(c []int32, k int32) int32 {
	energy := int32(100)
	lenCloud := len(c)
	currIdx := 0
	for {
		nextIdx := (currIdx + int(k)) % lenCloud
		currIdx = nextIdx

		// the current energy - next cloud value - 1
		// if the next cloud value 0 meaning it's only need - 1
		// if the next cloud value 1 meaning have to - 2 energy
		// that's why we add the - 1 at the end
		energy = energy - c[nextIdx] - 1
		if c[nextIdx] == 1 {
			energy -= 1
		}

		if nextIdx == 0 {
			return energy
		}

		if energy <= 0 {
			return energy
		}
	}
}

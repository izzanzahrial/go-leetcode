package stacknqueue

// https://leetcode.com/problems/asteroid-collision/?envType=study-plan-v2&envId=leetcode-75
func AsteroidCollision(asteroids []int) []int {
	var result []int

	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]

		if len(result) == 0 || result[len(result)-1] < 0 || asteroid > 0 {
			result = append(result, asteroid)
		} else {
			top := result[len(result)-1]

			if -asteroid == top {
				result = result[:len(result)-1]
			} else if -asteroid > top {
				result = result[:len(result)-1]
				i--
			}
		}
	}
	
	return result
}

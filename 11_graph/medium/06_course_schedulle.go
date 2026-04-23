package graph

// https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)

	for i := 0; i < len(prerequisites); i++ {
		// create course prerequisites map
		preMap[prerequisites[i][0]] = append(preMap[prerequisites[i][0]], prerequisites[i][1])
	}

	// visited map will be used to track per dfs route logic
	visited := make(map[int]bool)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := visited[course]; ok {
			return false
		}

		// if we already visit all of the prerequisites of the current course
		// we can finish the current course
		if preMap[course] == nil {
			return true
		}

		visited[course] = true
		// visit all the prerequisites of the current course
		for _, val := range preMap[course] {
			if bol := dfs(val); !bol {
				return false
			}
		}

		// since visited map only track the route of the current course
		// we need to remove the current course from the visited map
		// so if there are multiple routes to the current course
		// e.g. A -> B and C -> B, we need to remove the B from the visited map
		delete(visited, course)

		// the preMap also act a visited map but for the global route course
		// make sure all of the courses can be finished
		preMap[course] = []int{}

		return true
	}

	// check if all the courses can be finished
	// the courses start from 1
	for i := 1; i <= numCourses; i++ {
		if bol := dfs(i); !bol {
			return false
		}
	}

	return true
}

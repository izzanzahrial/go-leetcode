package graph

// https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)

	for i := 0; i < len(prerequisites); i++ {
		preMap[prerequisites[i][0]] = append(preMap[prerequisites[i][0]], prerequisites[i][1])
	}

	visited := make(map[int]bool)
	var dfs func(course int) bool

	dfs = func(course int) bool {
		if _, ok := visited[course]; ok {
			return false
		}

		if preMap[course] == nil {
			return true
		}

		visited[course] = true
		for _, val := range preMap[course] {
			if bol := dfs(val); !bol {
				return false
			}
		}
		delete(visited, course)
		preMap[course] = []int{}

		return true
	}

	for i := 1; i <= numCourses; i++ {
		if bol := dfs(i); !bol {
			return false
		}
	}

	return true
}

package hash

// https://leetcode.com/problems/keys-and-rooms/description/?envType=study-plan-v2&envId=leetcode-75
func CanVisitAllRooms(rooms [][]int) bool {
	visitedMap := make(map[int]bool, len(rooms))
	for i := 0; i < len(rooms); i++ {
		visitedMap[i] = false
	}
	visitedMap[0] = true

	// dfs check the room
	checkRoomKeys(rooms[0], rooms, visitedMap)

	if ok := checkVisited(visitedMap); !ok {
		return false
	}

	return true
}

func checkRoomKeys(keys []int, rooms [][]int, visited map[int]bool) {
	for _, key := range keys {
		if visited[key] {
			continue
		}

		visited[key] = true
		checkRoomKeys(rooms[key], rooms, visited)
	}
}

func checkVisited(visited map[int]bool) bool {
	for _, ok := range visited {
		if !ok {
			return false
		}
	}

	return true
}

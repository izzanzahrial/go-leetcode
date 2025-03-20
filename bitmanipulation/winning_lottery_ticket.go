package bitmanipulation

// https://www.hackerrank.com/challenges/winning-lottery-ticket/problem?isFullScreen=true
func winningLotteryTicket(tickets []string) int64 {
	var result int64
	// mapTickets := make(map[string]map[int]int, len(tickets)-1)

	// for i := 0; i < len(tickets); i++ {
	//     mapTickets[tickets[i]] = make(map[int]int, len(tickets[i]))
	//     for _, val := range tickets[i] {
	//         intVal,_ := strconv.Atoi(string(val))
	//         mapTickets[tickets[i]][intVal] += 1
	//     }
	// }

	// wg := &sync.WaitGroup{}
	// for i := 0; i < len(tickets)-1; i++ {
	// 	for j := i + 1; j < len(tickets); j++ {
	// first solution (slow)
	// currMap := make(map[int]int)
	// for _, val := range tickets[i] {
	//     intVal,_ := strconv.Atoi(string(val))
	//     currMap[intVal] += 1
	// }
	// for _, val := range tickets[j] {
	//     intVal,_ := strconv.Atoi(string(val))
	//     currMap[intVal] += 1
	// }

	// if checkMap(currMap) {
	//     result++
	// }

	// second solution (slow)
	// mapKeys := map[int]bool{
	//     0: true,
	//     1: true,
	//     2: true,
	//     3: true,
	//     4: true,
	//     5: true,
	//     6: true,
	//     7: true,
	//     8: true,
	//     9: true,
	// }
	// for _, val := range tickets[i] {
	//     intVal,_ := strconv.Atoi(string(val))
	//     if _, ok := mapKeys[intVal]; ok {
	//         delete(mapKeys, intVal)
	//     }
	// }
	// for _, val := range tickets[j] {
	//     intVal,_ := strconv.Atoi(string(val))
	//     if _, ok := mapKeys[intVal]; ok {
	//         delete(mapKeys, intVal)
	//     }
	// }

	// if len(mapKeys) == 0 {
	//     result++
	// }

	// second solution with go routine (still slow)
	// 		go func() {
	// 			defer func() {
	// 				wg.Done()
	// 			}()
	// 			wg.Add(1)
	// 			mapKeys := map[int]bool{
	// 				0: true,
	// 				1: true,
	// 				2: true,
	// 				3: true,
	// 				4: true,
	// 				5: true,
	// 				6: true,
	// 				7: true,
	// 				8: true,
	// 				9: true,
	// 			}
	// 			for _, val := range tickets[i] {
	// 				intVal, _ := strconv.Atoi(string(val))
	// 				if _, ok := mapKeys[intVal]; ok {
	// 					delete(mapKeys, intVal)
	// 				}
	// 			}
	// 			for _, val := range tickets[j] {
	// 				intVal, _ := strconv.Atoi(string(val))
	// 				if _, ok := mapKeys[intVal]; ok {
	// 					delete(mapKeys, intVal)
	// 				}
	// 			}

	// 			if len(mapKeys) == 0 {
	// 				result++
	// 			}
	// 		}()
	// 	}
	// }
	// wg.Wait()

	// third solution with bit manipulation (idk still slow)
	mapMasks := make(map[string]int, len(tickets))
	for _, ticket := range tickets {
		mask := 0
		for _, str := range ticket {
			// the ASCII of the str - the ASCII of '0'
			// e.g. '3' - '0' = 51 - 48 = 3
			digit := int(str - '0')
			mask |= 1 << digit
		}
		mapMasks[ticket] = mask
	}

	// because the bit manipulation of 1 shift left 10 times
	// equals to 1000000000, and minus 1 equals to
	// 0111111111
	goldenTicketMask := (1 << 10) - 1
	for i := 0; i < len(tickets)-1; i++ {
		for j := i + 1; j < len(tickets); j++ {
			// compare if the combination of mask ticket i and ticket j
			// equals to the golden ticket mask
			if goldenTicketMask == mapMasks[tickets[i]]|mapMasks[tickets[j]] {
				result++
			}
		}
	}

	return result
}

// func checkMap(data map[int]int) bool {
// 	for i := 0; i <= 9; i++ {
// 		if _, ok := data[i]; !ok {
// 			return false
// 		}
// 	}

// 	return true
// }

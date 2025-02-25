package sorting

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var result int
	for _, boxType := range boxTypes {
		// check if the number of box less than truckSize
		if boxType[0] < truckSize {
			truckSize -= boxType[0]
			result += boxType[0] * boxType[1]
		} else {
			// add the rest of the box by the value of the rest of the truckSize
			// because the truckSize already substract by the previous number of box
			result += truckSize * boxType[1]
			break
		}
	}

	return result
}

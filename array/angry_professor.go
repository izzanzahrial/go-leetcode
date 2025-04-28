package array

// https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true
func angryProfessor(k int32, a []int32) string {
	var arrivedStudents int
	for _, arrivalTime := range a {
		if arrivalTime <= 0 {
			arrivedStudents++
		}
	}

	// class is cancelled
	if arrivedStudents < int(k) {
		return "YES"
	}

	// class is not cancelled
	return "NO"
}

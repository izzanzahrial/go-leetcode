package array

// https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true
func GradingStudents(grades []int32) []int32 {
	for i, grade := range grades {
		if grade >= 38 {
			next5 := ((grade / 5) + 1) * 5
			if next5-grade < 3 {
				grades[i] = next5
			}
		}
	}

	return grades
}

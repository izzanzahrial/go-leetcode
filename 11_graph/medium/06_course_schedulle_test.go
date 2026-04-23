package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	prerequisites1 := [][]int{{1, 0}}
	prerequisites2 := [][]int{{1, 0}, {0, 1}}

	numCourses1 := 2
	numCourses2 := 2

	result1 := canFinish(numCourses1, prerequisites1)
	result2 := canFinish(numCourses2, prerequisites2)

	assert.Equal(t, result1, true)
	assert.Equal(t, result2, false)
}

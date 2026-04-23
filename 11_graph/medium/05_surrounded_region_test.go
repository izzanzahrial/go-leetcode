package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	subTest := []struct {
		name        string
		board       [][]byte
		expectedRes [][]byte
	}{
		{
			name:        "surrounded region many",
			board:       [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			expectedRes: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			name:        "surrounded region one",
			board:       [][]byte{{'X'}},
			expectedRes: [][]byte{{'X'}},
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			result := solve(test.board)
			assert.Equal(t, test.expectedRes, result)
		})
	}
}

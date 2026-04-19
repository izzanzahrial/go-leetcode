package backtrack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSearch(t *testing.T) {
	subTest := []struct {
		name        string
		board       [][]byte
		word        string
		expectedRes bool
	}{
		{
			name:        "word search true",
			board:       [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:        "ABCCED",
			expectedRes: true,
		},
		{
			name:        "word search false",
			board:       [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:        "ABCB",
			expectedRes: false,
		},
	}

	for _, test := range subTest {
		t.Run(test.name, func(t *testing.T) {
			res := wordSearch(test.board, test.word)
			assert.Equal(t, test.expectedRes, res)
		})
	}
}
